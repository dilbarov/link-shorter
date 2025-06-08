package link

import (
	"context"
	linkModels "link-shorter/internal/link/models"
	linkPayloads "link-shorter/internal/link/payloads"
	linkRepository "link-shorter/internal/link/repository"
)

const UpdateLinkCommandName = "link.create"

type UpdateCommand struct {
	Payload *linkPayloads.UpdatePayload
}

func (c UpdateCommand) Name() string {
	return UpdateLinkCommandName
}

type UpdateCommandHandler struct {
	LinkRepository linkRepository.Repository
}

func NewUpdateCommandHandler(repo linkRepository.Repository) *UpdateCommandHandler {
	return &UpdateCommandHandler{LinkRepository: repo}
}

func (h *UpdateCommandHandler) Handle(ctx context.Context, cmd UpdateCommand) (*linkModels.Model, error) {
	link, err := h.LinkRepository.GetById(cmd.Payload.Id)

	if err != nil {
		return nil, err
	}

	if cmd.Payload.Hash != nil && *cmd.Payload.Hash != link.Hash {
		link, err = h.LinkRepository.GetByHash(*cmd.Payload.Hash)
		if link != nil {
			return nil, err
		}
		link.Hash = *cmd.Payload.Hash
	}

	if cmd.Payload.Url != nil {
		link.Url = *cmd.Payload.Url
	}

	result, err := h.LinkRepository.Update(link)
	if err != nil {
		return nil, err
	}
	return result, nil
}
