package cqrs

import (
	"context"
)

// PublishEvent - обертка for EventBus.Publish
func PublishEvent(bus EventBus, ctx context.Context, event Event) []error {
	return bus.Publish(ctx, event)
}

// SubscribeToEvent - wrapper for  EventBus.Subscribe
// eventZeroVal is used to get reflect.Type
func SubscribeToEvent[E Event](bus EventBus, eventZeroVal E, handler EventHandler[E]) error {
	return bus.Subscribe(eventZeroVal, handler)
}
