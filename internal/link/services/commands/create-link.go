package link

import (
	"context"
	linkModels "link-shorter/internal/link/models"
	linkPayloads "link-shorter/internal/link/payloads"
	linkRepository "link-shorter/internal/link/repository"
)

const CreateLinkCommandName = "link.create"

type CreateCommand struct {
	Payload *linkPayloads.CreatePayload
}

func (c CreateCommand) Name() string {
	return CreateLinkCommandName
}

type CreateCommandHandler struct {
	LinkRepository linkRepository.Repository
}

func NewCreateCommandHandler(repo linkRepository.Repository) *CreateCommandHandler {
	return &CreateCommandHandler{LinkRepository: repo}
}

func (h *CreateCommandHandler) Handle(ctx context.Context, cmd CreateCommand) (*linkModels.Model, error) {
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
