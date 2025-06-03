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
