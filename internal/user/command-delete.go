package user

import "context"

type DeleteCommand struct {
	Payload *GetByIdParams
}

type DeleteCommandHandler struct {
	UserRepository Repository
}

func NewDeleteCommandHandler(repo Repository) *DeleteCommandHandler {
	return &DeleteCommandHandler{
		UserRepository: repo,
	}
}

func (h *DeleteCommandHandler) Handle(ctx context.Context, command DeleteCommand) error {
	err := h.UserRepository.Delete(command.Payload.Id)
	return err
}
