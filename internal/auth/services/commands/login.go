package auth

import (
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
	_, err := h.UserRepository.GetByEmail(cmd.Payload.Email)

	if err != nil {
		return "", err
	}

	return "", nil
}
