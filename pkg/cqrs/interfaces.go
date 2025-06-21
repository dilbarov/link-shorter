package cqrs

import "context"

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
