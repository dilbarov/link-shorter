package link

import (
	linkPayloads "link-shorter/internal/link/payloads"
	linkRepository "link-shorter/internal/link/repository"
)

type DeleteCommand struct {
	Payload *linkPayloads.GetByIDParams
}

type DeleteCommandHandler struct {
	LinkRepository linkRepository.Repository
}

func (h *DeleteCommandHandler) Execute(cmd DeleteCommand) error {
	err := h.LinkRepository.Delete(cmd.Payload.Id)

	if err != nil {
		return err
	}

	return nil
}
