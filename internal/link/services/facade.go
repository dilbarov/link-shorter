package link

import (
	linkRepository "link-shorter/internal/link/repository"
	linkCommands "link-shorter/internal/link/services/commands"
	linkQueries "link-shorter/internal/link/services/queries"
)

type ServiceFacade struct {
	Commands *CommandBus
	Queries  *QueryBus
}

func NewServiceFacade(repo linkRepository.Repository) *ServiceFacade {
	return &ServiceFacade{
		Commands: &CommandBus{
			Create: &linkCommands.CreateCommandHandler{
				LinkRepository: repo,
			},
			Update: &linkCommands.UpdateCommandHandler{
				LinkRepository: repo,
			},
			Delete: &linkCommands.DeleteCommandHandler{
				LinkRepository: repo,
			},
		},
		Queries: &QueryBus{
			GetByHash: &linkQueries.GetByHashQueryHandler{
				LinkRepository: repo,
			},
			GetById: &linkQueries.GetByIdQueryHandler{
				LinkRepository: repo,
			},
		},
	}
}
