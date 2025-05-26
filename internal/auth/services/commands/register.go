package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	authErrors "link-shorter/internal/auth/errors"
	authPayloads "link-shorter/internal/auth/payloads"
	userModels "link-shorter/internal/user/models"
	userRepository "link-shorter/internal/user/repository"
	"link-shorter/pkg/jwt"
)

type RegisterCommand struct {
	Payload authPayloads.RegisterRequest
}

type RegisterCommandHandler struct {
	UserRepository userRepository.Repository
	JwtService     *jwt.Service
}

func (h *RegisterCommandHandler) Execute(cmd RegisterCommand) (string, error) {
	existsUser, err := h.UserRepository.GetByEmail(cmd.Payload.Email)

	if existsUser != nil {
		return "", errors.New(authErrors.ErrEmailInUse)
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

	createdUser, err := h.UserRepository.Create(user)

	if err != nil {
		return "", err
	}

	token, err := h.JwtService.Create(createdUser.Id.String(), createdUser.Email)

	if err != nil {
		return "", err
	}

	return token, nil
}
