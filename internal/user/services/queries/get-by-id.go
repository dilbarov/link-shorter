package user

import (
	userModels "link-shorter/internal/user/models"
	userPayloads "link-shorter/internal/user/payloads"
	userRepository "link-shorter/internal/user/repository"
)

type GetByIdQuery struct {
	Params *userPayloads.GetByIdParams
}

type GetByIdQueryHandler struct {
	UserRepository userRepository.Repository
}

func (h *GetByIdQueryHandler) Execute(query GetByIdQuery) (*userModels.Model, error) {
	user, err := h.UserRepository.GetById(query.Params.Id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
