package user

import (
	userModels "link-shorter/internal/user/models"
	userPayloads "link-shorter/internal/user/payloads"
	userRepository "link-shorter/internal/user/repository"
)

type UpdateCommand struct {
	Payload *userPayloads.UpdatePayload
}

type UpdateCommandHandler struct {
	UserRepository userRepository.Repository
}

func (h *UpdateCommandHandler) Execute(cmd UpdateCommand) (*userModels.Model, error) {
	user, err := h.UserRepository.GetById(cmd.Payload.Id)

	if err != nil {
		return nil, err
	}

	user.Name = cmd.Payload.Name

	result, err := h.UserRepository.Update(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}
