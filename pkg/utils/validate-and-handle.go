package utils

import (
	"context"
	"fmt"
	"reflect"
)

func ValidateAndHandle(handlers map[reflect.Type]any, ctx context.Context, messageAsAny any) (result any, err error) {
	messageType := TypeOfValue(messageAsAny)

	handlerInterface, ok := handlers[messageType]
	if !ok {
		return nil, fmt.Errorf("no handler registered for message type %v", messageType)
	}

	// Using reflection to call the Handle method
	handlerValue := reflect.ValueOf(handlerInterface)
	queryValue := reflect.ValueOf(messageAsAny)

	method := handlerValue.MethodByName("Handle")

	if !method.IsValid() {
		return nil, fmt.Errorf("handler for %T does not have a Handle method", handlerInterface)
	}

	// Checking the types of arguments of the Handle method (context.Context, C)
	handleMethodType := method.Type()
	if handleMethodType.NumIn() != 2 {
		return nil, fmt.Errorf("handler Handle method for %T expects 2 arguments, got %d", handlerInterface, handleMethodType.NumIn())
	}
	// First argument must be context.Context
	// Second argument must be compatible with cmdAsAny type

	// Passing arguments
	args := []reflect.Value{
		reflect.ValueOf(ctx),
		queryValue,
	}

	results := method.Call(args) // []reflect.Value, expecting (R, error)

	if len(results) != 2 {
		return nil, fmt.Errorf("handler Handle method returned %d values, expected 2", len(results))
	}

	resultVal := results[0].Interface()
	errVal := results[1].Interface()

	if errVal == nil {
		return resultVal, nil
	}

	if err, ok := errVal.(error); ok {
		return resultVal, err // resultVal may not be nil even if there is an error
	}

	return resultVal, fmt.Errorf("handler returned non-error type for error value: %T", errVal)
}
