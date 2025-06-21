package auth

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"link-shorter/internal/user"
	"link-shorter/pkg/jwt"
)

type LoginCommand struct {
	Payload LoginRequest
}

type LoginCommandHandler struct {
	UserRepository user.Repository
	JwtService     *jwt.Service
}

func NewLoginCommandHandler(userRepository user.Repository, jwtService *jwt.Service) *LoginCommandHandler {
	return &LoginCommandHandler{
		UserRepository: userRepository,
		JwtService:     jwtService,
	}
}

func (h *LoginCommandHandler) Handle(ctx context.Context, cmd LoginCommand) (string, error) {
	existsUser, err := h.UserRepository.GetByEmail(cmd.Payload.Email)

	if err != nil {
		return "", err
	}

	if existsUser.PasswordHash == nil {
		return "", errors.New(ErrWrongCredentials)
	}

	err = bcrypt.CompareHashAndPassword([]byte(*existsUser.PasswordHash), []byte(cmd.Payload.Password))
	if err != nil {
		return "", errors.New(ErrWrongCredentials)
	}

	token, err := h.JwtService.Create(jwt.Data{
		Sub:   existsUser.Id.String(),
		Email: existsUser.Email,
	})

	if err != nil {
		return "", err
	}

	return token, nil
}
