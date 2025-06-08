package cqrs

import (
	"context"
	"fmt"
	cqrsInterfaces "link-shorter/pkg/cqrs/interfaces"
	"link-shorter/pkg/utils"
)

// ExecuteQuery - wrapper for QueryBus.ExecuteQuery
func ExecuteQuery[Q cqrsInterfaces.Query[R], R any](bus cqrsInterfaces.QueryBus, ctx context.Context, query Q) (R, error) {
	result, err := bus.Execute(ctx, query)
	var zeroR R

	if err != nil {
		return zeroR, err
	}

	// For queries, a result that is not a struct{} is usually always expected.
	if result == nil {
		// Check if type R can be nil (pointers, interfaces, slices, maps, channels, functions)
		if utils.CanBeNil[R]() {
			return zeroR, nil // zeroR is already nil, that's okay
		}
		return zeroR, fmt.Errorf("cqrs: ExecuteQuery: expected a result of type %T, but got nil", zeroR)
	}

	typedResult, ok := result.(R)
	if !ok {
		return zeroR, fmt.Errorf("cqrs: ExecuteQuery: unexpected result type: got %T, want %T", result, zeroR)
	}
	return typedResult, nil
}

// RegisterQueryHandler - wrapper for QueryBus.Register
// queryZeroVal is used to get reflect.Type
func RegisterQueryHandler[Q cqrsInterfaces.Query[R], R any](bus cqrsInterfaces.QueryBus, queryZeroVal Q, handler cqrsInterfaces.QueryHandler[Q, R]) {
	err := bus.Register(queryZeroVal, handler)

	if err != nil {
		panic(err)
	}
}
