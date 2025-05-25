package user

import (
	userPayloads "link-shorter/internal/user/payloads"
	userRepository "link-shorter/internal/user/repository"
)

type DeleteCommand struct {
	Payload *userPayloads.GetByIdParams
}

type DeleteCommandHandler struct {
	UserRepository userRepository.Repository
}

func (h *DeleteCommandHandler) Execute(command DeleteCommand) error {
	err := h.UserRepository.Delete(command.Payload.Id)
	return err
}
