package link

import (
	linkModels "link-shorter/internal/link/models"
	linkPayloads "link-shorter/internal/link/payloads"
	linkRepository "link-shorter/internal/link/repository"
)

type UpdateCommand struct {
	Payload *linkPayloads.UpdateRequest
}

type UpdateCommandHandler struct {
	LinkRepository linkRepository.Repository
}

func (h *UpdateCommandHandler) Execute(cmd UpdateCommand) (*linkModels.Model, error) {
	link := linkModels.NewLink(cmd.Payload.Url)
	result, err := h.LinkRepository.Update(cmd.Payload.Id, link)
	if err != nil {
		return nil, err
	}
	return result, nil
}
