package cqrs

import (
	"context"
	cqrsInterfaces "link-shorter/pkg/cqrs/interfaces"
)

// PublishEvent - обертка for EventBus.Publish
func PublishEvent(bus cqrsInterfaces.EventBus, ctx context.Context, event cqrsInterfaces.Event) []error {
	return bus.Publish(ctx, event)
}

// SubscribeToEvent - wrapper for  EventBus.Subscribe
// eventZeroVal is used to get reflect.Type
func SubscribeToEvent[E cqrsInterfaces.Event](bus cqrsInterfaces.EventBus, eventZeroVal E, handler cqrsInterfaces.EventHandler[E]) error {
	return bus.Subscribe(eventZeroVal, handler)
}
