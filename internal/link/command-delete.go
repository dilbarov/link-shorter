package link

import "context"

const DeleteLinkCommandName = "link.delete"

type DeleteCommand struct {
	Payload *GetByIDParams
}

func (c DeleteCommand) Name() string {
	return DeleteLinkCommandName
}

type DeleteCommandHandler struct {
	LinkRepository Repository
}

func NewDeleteCommandHandler(repo Repository) *DeleteCommandHandler {
	return &DeleteCommandHandler{LinkRepository: repo}
}

func (h *DeleteCommandHandler) Handle(ctx context.Context, cmd DeleteCommand) (bool, error) {
	err := h.LinkRepository.Delete(cmd.Payload.Id)
	if err != nil {
		return false, err
	}
	return true, nil
}
