package click

import (
	"context"
	clickModels "link-shorter/internal/click/models"
	clickPayloads "link-shorter/internal/click/payloads"
	clickRepository "link-shorter/internal/click/repository"
)

const AddClickCommandName = "click.add"

type AddCommand struct {
	Payload *clickPayloads.AddPayload
}

func (cmd AddCommand) Name() string {
	return AddClickCommandName
}

type AddCommandHandler struct {
	ClickRepository clickRepository.Repository
}

func NewAddCommandCommandHandler(clickRepository clickRepository.Repository) *AddCommandHandler {
	return &AddCommandHandler{
		ClickRepository: clickRepository,
	}
}

func (h *AddCommandHandler) Handle(ctx context.Context, cmd AddCommand) (any, error) {
	click := clickModels.NewClick(cmd.Payload.LinkId)
	_, err := h.ClickRepository.Create(click)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
