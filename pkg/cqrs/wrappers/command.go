package cqrs

import (
	"context"
	"fmt"
	cqrsInterfaces "link-shorter/pkg/cqrs/interfaces"
	"link-shorter/pkg/utils"
)

// ExecuteCommand - wrapper for CommandBus.ExecuteCommand
func ExecuteCommand[C cqrsInterfaces.Command[R], R any](bus cqrsInterfaces.CommandBus, ctx context.Context, cmd C) (R, error) {
	result, err := bus.Execute(ctx, cmd) // cmd is passed as any
	if err != nil {
		var zeroR R
		return zeroR, err
	}

	// If the command returns nothing (R is a struct{}), then result may be nil
	// or a special marker of success from the implementation of the bus.
	var typedResult R
	if utils.IsEmptyResult[R]() { // If R is a struct{}
		// Check if the bus implementation returned an error instead of nil for the result
		if result != nil {
			// Maybe the implementation returns an error for a "successful" void command,
			// or a special marker. It depends on the implementation contract.
			// For simplicity, if R is a struct{}, we simply return zeroR and err.
		}
		return typedResult, nil
	}

	if result == nil { // If a result is expected and nil is received
		// Check if type R can be nil (pointers, interfaces, slices, maps, channels, functions)
		if utils.CanBeNil[R]() {
			return typedResult, nil // typedResult is already nil, that's ok
		}
		return typedResult, fmt.Errorf("cqrs: ExecuteCommand: expected a result of type %T, but got nil", typedResult)
	}

	typedResult, ok := result.(R)
	if !ok {
		return typedResult, fmt.Errorf("cqrs: ExecuteCommand: unexpected result type: got %T, want %T", result, typedResult)
	}
	return typedResult, nil
}

// RegisterCommandHandler - wrapper for CommandBus.Register
// commandZeroVal is used to get reflect.Type
func RegisterCommandHandler[C cqrsInterfaces.Command[R], R any](bus cqrsInterfaces.CommandBus, commandZeroVal C, handler cqrsInterfaces.CommandHandler[C, R]) error {
	// We pass commandZeroVal, the bus implementation itself will take reflect.Type from it if necessary.
	return bus.Register(commandZeroVal, handler)
}
