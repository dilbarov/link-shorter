package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	authErrors "link-shorter/internal/auth/errors"
	authPayloads "link-shorter/internal/auth/payloads"
	userRepository "link-shorter/internal/user/repository"
)

type LoginCommand struct {
	Payload authPayloads.LoginRequest
}

type LoginCommandHandler struct {
	UserRepository userRepository.Repository
}

func (h *LoginCommandHandler) Execute(cmd LoginCommand) (string, error) {
	existsUser, err := h.UserRepository.GetByEmail(cmd.Payload.Email)

	if err != nil {
		return "", err
	}

	if existsUser.PasswordHash == nil {
		return "", errors.New(authErrors.ErrWrongCredentials)
	}

	err = bcrypt.CompareHashAndPassword([]byte(*existsUser.PasswordHash), []byte(cmd.Payload.Password))
	if err != nil {
		return "", errors.New(authErrors.ErrWrongCredentials)
	}

	return "", nil
}
