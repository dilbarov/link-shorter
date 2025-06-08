package cqrs

import cqrsInterfaces "link-shorter/pkg/cqrs/interfaces"

func InitInMemory() (cqrsInterfaces.CommandBus, cqrsInterfaces.QueryBus, cqrsInterfaces.EventBus) {
	return NewInMemoryCommandBus(), NewInMemoryQueryBus(), NewInMemoryEventBus()
}
