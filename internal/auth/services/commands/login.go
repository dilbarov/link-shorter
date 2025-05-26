package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	authErrors "link-shorter/internal/auth/errors"
	authPayloads "link-shorter/internal/auth/payloads"
	userRepository "link-shorter/internal/user/repository"
	"link-shorter/pkg/jwt"
)

type LoginCommand struct {
	Payload authPayloads.LoginRequest
}

type LoginCommandHandler struct {
	UserRepository userRepository.Repository
	JwtService     *jwt.Service
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

	token, err := h.JwtService.Create(existsUser.Id.String(), existsUser.Email)

	if err != nil {
		return "", err
	}

	return token, nil
}
