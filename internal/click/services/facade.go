package click

import (
	"context"
	clickRepository "link-shorter/internal/click/repository"
	clickResponses "link-shorter/internal/click/responses"
	clickCommands "link-shorter/internal/click/services/commands"
	clickQueries "link-shorter/internal/click/services/queries"
	cqrsInterfaces "link-shorter/pkg/cqrs/interfaces"
	cqrsWrappers "link-shorter/pkg/cqrs/wrappers"
)

type commands struct {
	Add func(ctx context.Context, cmd clickCommands.AddCommand) (any, error)
}

type queries struct {
	GetCountByLink func(ctx context.Context, cmd clickQueries.GetCountQuery) (int64, error)
	GetStatsByLink func(ctx context.Context, cmd clickQueries.GetStatsQuery) ([]*clickResponses.CountStatItem, error)
}

type ServiceFacadeDeps struct {
	CommandBus      cqrsInterfaces.CommandBus
	QueryBus        cqrsInterfaces.QueryBus
	EventBus        cqrsInterfaces.EventBus
	ClickRepository clickRepository.Repository
}

type ServiceFacade struct {
	Commands *commands
	Queries  *queries
}

func NewServiceFacade(deps ServiceFacadeDeps) *ServiceFacade {
	// Commands
	cqrsWrappers.RegisterCommandHandler[clickCommands.AddCommand, any](
		deps.CommandBus,
		clickCommands.AddCommand{},
		clickCommands.NewAddCommandCommandHandler(deps.ClickRepository),
	)

	// Queries
	cqrsWrappers.RegisterQueryHandler[clickQueries.GetCountQuery, int64](
		deps.QueryBus,
		clickQueries.GetCountQuery{},
		clickQueries.NewGetCountQueryHandler(deps.ClickRepository),
	)

	cqrsWrappers.RegisterQueryHandler[clickQueries.GetStatsQuery, []*clickResponses.CountStatItem](
		deps.QueryBus,
		clickQueries.GetStatsQuery{},
		clickQueries.NewGetStatsQueryHandler(deps.ClickRepository),
	)

	return &ServiceFacade{
		Commands: &commands{
			Add: func(ctx context.Context, cmd clickCommands.AddCommand) (any, error) {
				return cqrsWrappers.ExecuteCommand[clickCommands.AddCommand, any](deps.CommandBus, ctx, cmd)
			},
		},
		Queries: &queries{
			GetCountByLink: func(ctx context.Context, query clickQueries.GetCountQuery) (int64, error) {
				return cqrsWrappers.ExecuteQuery[clickQueries.GetCountQuery, int64](deps.QueryBus, ctx, query)
			},
			GetStatsByLink: func(ctx context.Context, query clickQueries.GetStatsQuery) ([]*clickResponses.CountStatItem, error) {
				return cqrsWrappers.ExecuteQuery[clickQueries.GetStatsQuery, []*clickResponses.CountStatItem](deps.QueryBus, ctx, query)
			},
		},
	}
}
