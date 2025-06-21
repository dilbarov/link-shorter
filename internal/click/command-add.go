package click

import (
	"context"
)

type AddCommand struct {
	Payload AddPayload
}

type AddCommandHandler struct {
	ClickRepository Repository
}

func NewAddCommandCommandHandler(clickRepository Repository) *AddCommandHandler {
	return &AddCommandHandler{
		ClickRepository: clickRepository,
	}
}

func (h *AddCommandHandler) Handle(ctx context.Context, cmd AddCommand) (any, error) {
	click := NewClick(cmd.Payload.LinkId)
	_, err := h.ClickRepository.Create(click)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
