package link

import (
	"context"
)

type ServiceFacadeDeps struct {
	LinkRepository Repository
}

type ServiceFacade struct {
	Create    func(ctx context.Context, cmd CreateCommand) (*Model, error)
	Update    func(ctx context.Context, cmd UpdateCommand) (*Model, error)
	Delete    func(ctx context.Context, cmd DeleteCommand) (bool, error)
	GetByHash func(ctx context.Context, q GetByHashQuery) (*Model, error)
	GetById   func(ctx context.Context, q GetByIdQuery) (*Model, error)
	GetAll    func(ctx context.Context, q GetAllQuery) ([]*Model, int, error)
}

func NewServiceFacade(deps ServiceFacadeDeps) *ServiceFacade {
	// Register Commands
	createHandler := NewCreateCommandHandler(deps.LinkRepository)
	updateHandler := NewUpdateCommandHandler(deps.LinkRepository)
	deleteHandler := NewDeleteCommandHandler(deps.LinkRepository)

	// Register Queries
	getByHashHandler := NewGetByHashQueryHandler(deps.LinkRepository)
	getByIdHandler := NewGetByIdQueryHandler(deps.LinkRepository)
	getAllHandler := NewGetAllQueryHandler(deps.LinkRepository)

	return &ServiceFacade{
		Create: createHandler.Handle,
		Update: updateHandler.Handle,
		Delete: deleteHandler.Handle,

		GetByHash: getByHashHandler.Handle,
		GetById:   getByIdHandler.Handle,
		GetAll:    getAllHandler.Handle,
	}
}
