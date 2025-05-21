package link

import (
	linkModels "link-shorter/internal/link/models"
	linkPayloads "link-shorter/internal/link/payloads"
	linkRepository "link-shorter/internal/link/repository"
)

type CreateCommand struct {
	Payload *linkPayloads.CreateRequest
}

type CreateCommandHandler struct {
	LinkRepository linkRepository.Repository
}

func (h *CreateCommandHandler) Execute(cmd CreateCommand) (*linkModels.Model, error) {
	link := linkModels.NewLink(cmd.Payload.Url)
	createdLink, err := h.LinkRepository.Create(link)
	if err != nil {
		return nil, err
	}
	return createdLink, nil
}
