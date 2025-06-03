package cqrs

import "context"

type Query[R any] interface {
	Name() string
}

type QueryHandler[Q Query[R], R any] interface {
	Handle(ctx context.Context, query Q) (R, error)
}

type QueryBus interface {
	Execute(ctx context.Context, queryAsAny any) (resultAsAny any, err error)
	Register(queryTypeOrZeroVal any, handlerAsAny any) error
}
