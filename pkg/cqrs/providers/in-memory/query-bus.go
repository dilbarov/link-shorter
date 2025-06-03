package cqrs

import (
	"context"
	"fmt"
	"link-shorter/pkg/utils"
	"reflect"
	"sync"
)

type InMemoryQueryBus struct {
	handlers map[reflect.Type]any
	mu       sync.RWMutex
}

func NewInMemoryQueryBus() *InMemoryQueryBus {
	return &InMemoryQueryBus{handlers: make(map[reflect.Type]any)}
}

func (b *InMemoryQueryBus) Register(queryTypeOrZeroVal any, handlerAsAny any) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	rType := utils.ResolveType(queryTypeOrZeroVal)

	if _, ok := b.handlers[rType]; ok {
		return fmt.Errorf("handler for query type %v already registered in InMemoryQueryBus", rType)
	}

	b.handlers[rType] = handlerAsAny
	return nil
}

func (b *InMemoryQueryBus) Execute(ctx context.Context, queryAsAny any) (any, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return utils.ValidateAndHandle(b.handlers, ctx, queryAsAny)
}
