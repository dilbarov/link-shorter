package link

import (
	linkRepository "link-shorter/internal/link/repository"
	linkCommands "link-shorter/internal/link/services/commands"
	"link-shorter/internal/link/services/queries"
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
		},
		Queries: &QueryBus{
			GetByHash: &queries.GetByHashQueryHandler{
				LinkRepository: repo,
			},
			GetById: &queries.GetByIdQueryHandler{
				LinkRepository: repo,
			},
		},
	}
}
