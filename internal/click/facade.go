package click

import (
	"context"
	"link-shorter/pkg/cqrs"
)

type ServiceFacadeDeps struct {
	EventBus        cqrs.EventBus
	ClickRepository Repository
}

type ServiceFacade struct {
	Add            func(ctx context.Context, cmd AddCommand) (any, error)
	GetCountByLink func(ctx context.Context, cmd GetCountQuery) (int64, error)
	GetStatsByLink func(ctx context.Context, cmd GetStatsQuery) ([]*CountStatItem, error)
}

func NewServiceFacade(deps ServiceFacadeDeps) *ServiceFacade {
	// Commands
	addHandler := NewAddCommandCommandHandler(deps.ClickRepository)

	// Queries
	getCountHandler := NewGetCountQueryHandler(deps.ClickRepository)
	getStatsHandler := NewGetStatsQueryHandler(deps.ClickRepository)

	return &ServiceFacade{
		Add:            addHandler.Handle,
		GetCountByLink: getCountHandler.Handle,
		GetStatsByLink: getStatsHandler.Handle,
	}
}
