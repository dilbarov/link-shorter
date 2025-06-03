package cqrs

import (
	"context"
	"fmt"
	cqrsInterfaces "link-shorter/pkg/cqrs/interfaces"
	"link-shorter/pkg/utils"
	"reflect"
	"sync"
)

type InMemoryEventBus struct {
	handlers map[reflect.Type][]any
	mu       sync.RWMutex
}

func NewInMemoryEventBus() *InMemoryEventBus {
	return &InMemoryEventBus{
		handlers: make(map[reflect.Type][]any),
	}
}

func (b *InMemoryEventBus) Subscribe(eventTypeOrZeroVal any, handlerAsAny any) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	rType := utils.ResolveType(eventTypeOrZeroVal)

	// Add a handler to the slice for this event type
	b.handlers[rType] = append(b.handlers[rType], handlerAsAny)
	return nil
}

func (b *InMemoryEventBus) Publish(ctx context.Context, event cqrsInterfaces.Event) []error {
	b.mu.RLock()

	eventType := utils.TypeOfValue(event)

	specificEventHandlers, ok := b.handlers[eventType]
	if !ok || len(specificEventHandlers) == 0 {
		b.mu.RUnlock()
		return nil
	}

	// Copy the slice of handlers so that the mutex can be unlocked early.
	// This allows other operations (like Subscribe) to not have to wait for all handlers to complete.
	handlersToCall := make([]any, len(specificEventHandlers))
	copy(handlersToCall, specificEventHandlers)
	b.mu.RUnlock()

	var wg sync.WaitGroup
	// Channel for collecting errors from handlers. Buffer size equals to number of handlers,
	// so that sending to the channel does not block handler goroutines if `Publish`
	// does not read from it fast enough (although in this case we read everything at once).
	errChan := make(chan error, len(handlersToCall))
	eventValue := reflect.ValueOf(event)

	for _, handlerInterface := range handlersToCall {
		wg.Add(1)

		go func(eventHandler any) {
			defer wg.Done()
			defer func() {
				if panicValue := recover(); panicValue != nil {
					// Sending a panic error to the error channel
					errChan <- fmt.Errorf("panic in event eventHandler %T for event %T: %v", eventHandler, event, panicValue)
				}
			}()

			handlerValue := reflect.ValueOf(eventHandler)
			method := handlerValue.MethodByName("Handle")

			if !method.IsValid() {
				errChan <- fmt.Errorf("eventHandler %T for event %T does not have a valid Handle method", eventHandler, event)
				return
			}

			// Passing arguments (context.Context, Event E)
			args := []reflect.Value{
				reflect.ValueOf(ctx),
				eventValue,
			}

			results := method.Call(args) // Expecting (error)

			if len(results) != 1 {
				errChan <- fmt.Errorf("event eventHandler %T Handle method for event %T returned %d values, expected 1 (error)", eventHandler, event, len(results))
				return
			}

			errVal := results[0].Interface()
			if errVal != nil {
				if err, ok := errVal.(error); ok {
					errChan <- fmt.Errorf("error from eventHandler %T for event %T: %w", eventHandler, event, err)
				} else {
					errChan <- fmt.Errorf("event eventHandler %T for event %T returned non-error type for error value: %T", eventHandler, event, errVal)
				}
			}
			// If there is no error, we do not send anything to errChan
		}(handlerInterface) // Passing copy of variables to the goroutine closure
	}

	// Goroutine to close the error channel after all handlers have completed.
	// This is important so that the loop reading from the channel below doesn't block forever.
	go func() {
		wg.Wait()
		close(errChan)
	}()

	var errs []error
	for err := range errChan {
		errs = append(errs, err)
	}

	if len(errs) == 0 {
		return nil
	}
	return errs
}
