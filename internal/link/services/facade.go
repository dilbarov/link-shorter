package link

import (
	"context"
	linkModels "link-shorter/internal/link/models"
	linkRepository "link-shorter/internal/link/repository"
	linkCommands "link-shorter/internal/link/services/commands"
	linkQueries "link-shorter/internal/link/services/queries"
	cqrsInterfaces "link-shorter/pkg/cqrs/interfaces"
	cqrsWrappers "link-shorter/pkg/cqrs/wrappers"
	"link-shorter/pkg/db"
)

type ServiceFacadeDeps struct {
	CommandBus     cqrsInterfaces.CommandBus
	QueryBus       cqrsInterfaces.QueryBus
	EventBus       cqrsInterfaces.EventBus
	LinkRepository linkRepository.Repository
}

type commands struct {
	Create func(ctx context.Context, cmd linkCommands.CreateCommand) (*linkModels.Model, error)
	Update func(ctx context.Context, cmd linkCommands.UpdateCommand) (*linkModels.Model, error)
	Delete func(ctx context.Context, cmd linkCommands.DeleteCommand) (bool, error)
}

type queries struct {
	GetByHash func(ctx context.Context, q linkQueries.GetByHashQuery) (*linkModels.Model, error)
	GetById   func(ctx context.Context, q linkQueries.GetByIdQuery) (*linkModels.Model, error)
	GetAll    func(ctx context.Context, q linkQueries.GetAllQuery) (*db.ListResult[*linkModels.Model], error)
}

type ServiceFacade struct {
	Commands *commands
	Queries  *queries
}

func NewServiceFacade(deps ServiceFacadeDeps) *ServiceFacade {
	// Register Commands
	cqrsWrappers.RegisterCommandHandler[linkCommands.CreateCommand, *linkModels.Model](
		deps.CommandBus,
		linkCommands.CreateCommand{},
		linkCommands.NewCreateCommandHandler(deps.LinkRepository),
	)

	cqrsWrappers.RegisterCommandHandler[linkCommands.DeleteCommand, bool](
		deps.CommandBus,
		linkCommands.DeleteCommand{},
		linkCommands.NewDeleteCommandHandler(deps.LinkRepository),
	)

	cqrsWrappers.RegisterCommandHandler[linkCommands.UpdateCommand, *linkModels.Model](
		deps.CommandBus,
		linkCommands.UpdateCommand{},
		linkCommands.NewUpdateCommandHandler(deps.LinkRepository),
	)

	// Register Queries
	cqrsWrappers.RegisterQueryHandler[linkQueries.GetByHashQuery, *linkModels.Model](
		deps.QueryBus,
		linkQueries.GetByHashQuery{},
		linkQueries.NewGetByHashQueryHandler(deps.LinkRepository),
	)

	cqrsWrappers.RegisterQueryHandler[linkQueries.GetByIdQuery, *linkModels.Model](
		deps.QueryBus,
		linkQueries.GetByIdQuery{},
		linkQueries.NewGetByIdQueryHandler(deps.LinkRepository),
	)

	cqrsWrappers.RegisterQueryHandler[linkQueries.GetAllQuery, *db.ListResult[*linkModels.Model]](
		deps.QueryBus,
		linkQueries.GetAllQuery{},
		linkQueries.NewGetAllQueryHandler(deps.LinkRepository),
	)

	return &ServiceFacade{
		Commands: &commands{
			Create: func(ctx context.Context, cmd linkCommands.CreateCommand) (*linkModels.Model, error) {
				return cqrsWrappers.ExecuteCommand[linkCommands.CreateCommand, *linkModels.Model](deps.CommandBus, ctx, cmd)
			},
			Update: func(ctx context.Context, cmd linkCommands.UpdateCommand) (*linkModels.Model, error) {
				return cqrsWrappers.ExecuteCommand[linkCommands.UpdateCommand, *linkModels.Model](deps.CommandBus, ctx, cmd)
			},
			Delete: func(ctx context.Context, cmd linkCommands.DeleteCommand) (bool, error) {
				return cqrsWrappers.ExecuteCommand[linkCommands.DeleteCommand, bool](deps.CommandBus, ctx, cmd)
			},
		},
		Queries: &queries{
			GetByHash: func(ctx context.Context, q linkQueries.GetByHashQuery) (*linkModels.Model, error) {
				return cqrsWrappers.ExecuteQuery[linkQueries.GetByHashQuery, *linkModels.Model](deps.QueryBus, ctx, q)
			},
			GetById: func(ctx context.Context, q linkQueries.GetByIdQuery) (*linkModels.Model, error) {
				return cqrsWrappers.ExecuteQuery[linkQueries.GetByIdQuery, *linkModels.Model](deps.QueryBus, ctx, q)
			},
			GetAll: func(ctx context.Context, q linkQueries.GetAllQuery) (*db.ListResult[*linkModels.Model], error) {
				return cqrsWrappers.ExecuteQuery[linkQueries.GetAllQuery, *db.ListResult[*linkModels.Model]](deps.QueryBus, ctx, q)
			},
		},
	}
}
