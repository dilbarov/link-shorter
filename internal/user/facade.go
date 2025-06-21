package user

import "context"

type ServiceFacade struct {
	Create     func(ctx context.Context, cmd CreateCommand) (*Model, error)
	Update     func(ctx context.Context, cmd UpdateCommand) (*Model, error)
	Delete     func(ctx context.Context, cmd DeleteCommand) error
	GetById    func(ctx context.Context, q GetByIdQuery) (*Model, error)
	GetByEmail func(ctx context.Context, q GetByEmailQuery) (*Model, error)
	GetAll     func(ctx context.Context, q GetAllQuery) ([]*Model, int, error)
}

func NewServiceFacade(repo Repository) *ServiceFacade {
	// Register Commands
	createHandler := NewCreateCommandHandler(repo)
	updateHandler := NewUpdateCommandHandler(repo)
	deleteHandler := NewDeleteCommandHandler(repo)

	// Register Queries
	getAllHandler := NewGetAllQueryHandler(repo)
	getByEmailHandler := NewGetByEmailQueryHandler(repo)
	getByIdHandler := NewGetByIdQueryHandler(repo)

	return &ServiceFacade{
		Create:     createHandler.Handle,
		Update:     updateHandler.Handle,
		Delete:     deleteHandler.Handle,
		GetById:    getByIdHandler.Handle,
		GetByEmail: getByEmailHandler.Handle,
		GetAll:     getAllHandler.Handle,
	}
}
