package auth

import (
	"errors"
	authErrors "link-shorter/internal/auth/errors"
	authPayloads "link-shorter/internal/auth/payloads"
	userModels "link-shorter/internal/user/models"
	userRepository "link-shorter/internal/user/repository"
)

type RegisterCommand struct {
	Payload authPayloads.RegisterRequest
}

type RegisterCommandHandler struct {
	UserRepository userRepository.Repository
}

func (h *RegisterCommandHandler) Execute(cmd RegisterCommand) (string, error) {
	existsUser, err := h.UserRepository.GetByEmail(cmd.Payload.Email)

	if existsUser != nil {
		return "", errors.New(authErrors.ErrEmailInUse)
	}

	if err != nil {
		return "", err
	}

	user := &userModels.Model{
		Email:        cmd.Payload.Email,
		PasswordHash: nil,
		Name:         &cmd.Payload.Name,
	}

	_, err = h.UserRepository.Create(user)

	if err != nil {
		return "", err
	}

	return user.Email, nil
}
