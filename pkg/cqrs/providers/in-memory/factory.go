package cqrs

func InitInMemory() (*InMemoryCommandBus, *InMemoryQueryBus, *InMemoryEventBus) {
	return NewInMemoryCommandBus(), NewInMemoryQueryBus(), NewInMemoryEventBus()
}
