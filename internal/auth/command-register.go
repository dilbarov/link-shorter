package auth

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"link-shorter/internal/user"
	"link-shorter/pkg/jwt"
)

type RegisterCommand struct {
	Payload RegisterRequest
}

type RegisterCommandHandler struct {
	UserRepository user.Repository
	JwtService     *jwt.Service
}

func NewRegisterCommandHandler(jwtService *jwt.Service, userRepository user.Repository) *RegisterCommandHandler {
	return &RegisterCommandHandler{
		UserRepository: userRepository,
		JwtService:     jwtService,
	}
}

func (h *RegisterCommandHandler) Handle(ctx context.Context, cmd RegisterCommand) (string, error) {
	existsUser, err := h.UserRepository.GetByEmail(cmd.Payload.Email)

	if existsUser != nil {
		return "", errors.New(ErrEmailInUse)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	passwordHash := string(hashedPassword)

	userModel := &user.Model{
		Email:        cmd.Payload.Email,
		PasswordHash: &passwordHash,
		Name:         &cmd.Payload.Name,
	}

	createdUser, err := h.UserRepository.Create(userModel)

	if err != nil {
		return "", err
	}

	token, err := h.JwtService.Create(jwt.Data{
		Sub:   createdUser.Id.String(),
		Email: createdUser.Email,
	})

	if err != nil {
		return "", err
	}

	return token, nil
}
