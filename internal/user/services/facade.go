package user

import (
	userRepository "link-shorter/internal/user/repository"
	userCommands "link-shorter/internal/user/services/commands"
	userQueries "link-shorter/internal/user/services/queries"
)

type ServiceFacade struct {
	Commands *CommandBus
	Queries  *QueryBus
}

func NewServiceFacade(repo userRepository.Repository) *ServiceFacade {
	return &ServiceFacade{
		Commands: &CommandBus{
			Create: &userCommands.CreateCommandHandler{
				UserRepository: repo,
			},
			Update: &userCommands.UpdateCommandHandler{
				UserRepository: repo,
			},
			Delete: &userCommands.DeleteCommandHandler{
				UserRepository: repo,
			},
		},
		Queries: &QueryBus{
			GetById: &userQueries.GetByIdQueryHandler{
				UserRepository: repo,
			},
			GetByEmail: &userQueries.GetByEmailQueryHandler{
				UserRepository: repo,
			},
			GetAll: &userQueries.GetAllQueryHandler{
				UserRepository: repo,
			},
		},
	}
}
