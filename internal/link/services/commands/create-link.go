package link

import (
	linkModels "link-shorter/internal/link/models"
	linkPayloads "link-shorter/internal/link/payloads"
	linkRepository "link-shorter/internal/link/repository"
)

type CreateCommand struct {
	Payload *linkPayloads.CreatePayload
}

type CreateCommandHandler struct {
	LinkRepository linkRepository.Repository
}

func (h *CreateCommandHandler) Execute(cmd CreateCommand) (*linkModels.Model, error) {
	link := linkModels.NewLink(cmd.Payload.Url)
	for {
		existsLink, _ := h.LinkRepository.GetByHash(link.Hash)
		if existsLink == nil {
			break
		}
		link.GenerateHash()
	}
	createdLink, err := h.LinkRepository.Create(link)
	if err != nil {
		return nil, err
	}
	return createdLink, nil
}
