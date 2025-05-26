package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
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

	if cmd.Payload.Password != "" {

	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	passwordHash := string(hashedPassword)

	user := &userModels.Model{
		Email:        cmd.Payload.Email,
		PasswordHash: &passwordHash,
		Name:         &cmd.Payload.Name,
	}

	_, err = h.UserRepository.Create(user)

	if err != nil {
		return "", err
	}

	return user.Email, nil
}
