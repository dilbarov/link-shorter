package cqrs

import "context"

type Command[R any] interface {
	Name() string
}

type CommandHandler[C Command[R], R any] interface {
	Handle(ctx context.Context, cmd C) (R, error)
}

type CommandBus interface {
	Execute(ctx context.Context, cmdAsAny any) (resultAsAny any, err error)
	Register(commandTypeOrZeroVal any, handlerAsAny any) error
}

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

type Event interface {
	Name() string
}

type EventHandler[E Event] interface {
	Handle(ctx context.Context, event E) error
}

type EventBus interface {
	Publish(ctx context.Context, eventAsAny any) []error
	Subscribe(eventTypeOrZeroVal any, handlerAsAny any) error
}
