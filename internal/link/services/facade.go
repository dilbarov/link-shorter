package link

import (
	"link-shorter/internal/link/repository"
	commands "link-shorter/internal/link/services/commands"
	"link-shorter/internal/link/services/queries"
)

type ServiceFacade struct {
	Commands *CommandBus
	Queries  *QueryBus
}

func NewServiceFacade(repo repository.LinkRepository) *ServiceFacade {
	return &ServiceFacade{
		Commands: &CommandBus{
			Create: &commands.CreateCommandHandler{
				LinkRepository: repo,
			},
			Update: &commands.UpdateCommandHandler{
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
