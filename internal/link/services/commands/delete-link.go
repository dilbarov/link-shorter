package link

import (
	"context"
	linkPayloads "link-shorter/internal/link/payloads"
	linkRepository "link-shorter/internal/link/repository"
)

const DeleteLinkCommandName = "link.delete"

type DeleteCommand struct {
	Payload *linkPayloads.GetByIDParams
}

func (c DeleteCommand) Name() string {
	return DeleteLinkCommandName
}

type DeleteCommandHandler struct {
	LinkRepository linkRepository.Repository
}

func NewDeleteCommandHandler(repo linkRepository.Repository) *DeleteCommandHandler {
	return &DeleteCommandHandler{LinkRepository: repo}
}

func (h *DeleteCommandHandler) Handle(ctx context.Context, cmd DeleteCommand) (bool, error) {
	err := h.LinkRepository.Delete(cmd.Payload.Id)
	if err != nil {
		return false, err
	}
	return true, nil
}
