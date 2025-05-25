package user

import (
	"errors"
	userErrors "link-shorter/internal/user/errors"
	userModels "link-shorter/internal/user/models"
	userPayloads "link-shorter/internal/user/payloads"
	userRepository "link-shorter/internal/user/repository"
)

type CreateCommand struct {
	Payload *userPayloads.CreatePayload
}

type CreateCommandHandler struct {
	UserRepository userRepository.Repository
}

func (h *CreateCommandHandler) Execute(cmd CreateCommand) (*userModels.Model, error) {
	user := userModels.NewUser(cmd.Payload.Email, &cmd.Payload.Password, &cmd.Payload.Name)

	existsUser, err := h.UserRepository.GetByEmail(user.Email)

	if existsUser != nil {
		return nil, errors.New(userErrors.ErrUserExists)
	}

	createdUser, err := h.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
