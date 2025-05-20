package link

import (
	"link-shorter/internal/link/repository"
	commands "link-shorter/internal/link/services/commands"
)

type ServiceFacade struct {
	Commands *CommandBus
	Queries  *QueryBus
}

func NewServiceFacade(repo repository.LinkRepository) *ServiceFacade {
	return &ServiceFacade{
		Commands: &CommandBus{
			CreateHandler: &commands.CreateHandler{
				Repo: repo,
			},
			UpdateHandler: &commands.UpdateHandler{
				Repo: repo,
			},
		},
		Queries: &QueryBus{},
	}
}
