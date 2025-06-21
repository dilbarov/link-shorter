package user

import (
	"context"
	"errors"
)

type CreateCommand struct {
	Payload *CreatePayload
}

type CreateCommandHandler struct {
	UserRepository Repository
}

func NewCreateCommandHandler(repo Repository) *CreateCommandHandler {
	return &CreateCommandHandler{
		UserRepository: repo,
	}
}

func (h *CreateCommandHandler) Handle(ctx context.Context, cmd CreateCommand) (*Model, error) {
	user := NewUser(cmd.Payload.Email, &cmd.Payload.Password, &cmd.Payload.Name)

	existsUser, err := h.UserRepository.GetByEmail(user.Email)

	if existsUser != nil {
		return nil, errors.New(ErrUserExists)
	}

	createdUser, err := h.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
