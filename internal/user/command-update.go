package user

import "context"

type UpdateCommand struct {
	Payload *UpdatePayload
}

type UpdateCommandHandler struct {
	UserRepository Repository
}

func NewUpdateCommandHandler(repo Repository) *UpdateCommandHandler {
	return &UpdateCommandHandler{
		UserRepository: repo,
	}
}

func (h *UpdateCommandHandler) Handle(ctx context.Context, cmd UpdateCommand) (*Model, error) {
	user, err := h.UserRepository.GetById(cmd.Payload.Id)

	if err != nil {
		return nil, err
	}

	user.Name = cmd.Payload.Name

	result, err := h.UserRepository.Update(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}
