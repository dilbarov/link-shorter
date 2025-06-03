package cqrs

import (
	"context"
	"fmt"
	"link-shorter/pkg/utils"
	"reflect"
	"sync"
)

type InMemoryCommandBus struct {
	handlers map[reflect.Type]any
	mu       sync.RWMutex
}

func NewInMemoryCommandBus() *InMemoryCommandBus {
	return &InMemoryCommandBus{
		handlers: make(map[reflect.Type]any),
	}
}

// Register is an InMemoryCommandBus method intended for registering a
// handler (handlerAsAny) for a particular command type (commandTypeOrZeroVal).
func (b *InMemoryCommandBus) Register(commandTypeOrZeroVal any, handlerAsAny any) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	rType := utils.ResolveType(commandTypeOrZeroVal)

	/**
	TODO:
		Validation that handlerAsAny is indeed a CommandHandler
		and that commandTypeOrZeroVal is indeed a Command can be added here via reflection,
		checking whether handlerAsAny implements the CommandHandler interface for type rType.
		This is complicated, but improves reliability. We'll skip it for brevity.
	*/

	if _, ok := b.handlers[rType]; ok {
		return fmt.Errorf("handler for command type %v already registered in InMemoryCommandBus", rType)
	}

	b.handlers[rType] = handlerAsAny
	return nil
}

func (b *InMemoryCommandBus) Execute(ctx context.Context, cmdAsAny any) (any, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return utils.ValidateAndHandle(b.handlers, ctx, cmdAsAny)
}
