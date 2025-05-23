package link

import (
	linkModels "link-shorter/internal/link/models"
	linkPayloads "link-shorter/internal/link/payloads"
	linkRepository "link-shorter/internal/link/repository"
)

type UpdateCommand struct {
	Payload *linkPayloads.UpdatePayload
}

type UpdateCommandHandler struct {
	LinkRepository linkRepository.Repository
}

func (h *UpdateCommandHandler) Execute(cmd UpdateCommand) (*linkModels.Model, error) {
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
