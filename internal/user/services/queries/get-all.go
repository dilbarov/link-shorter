package user

import (
	userModels "link-shorter/internal/user/models"
	userPayloads "link-shorter/internal/user/payloads"
	userRepository "link-shorter/internal/user/repository"
)

type GetAllQuery struct {
	Params *userPayloads.GetAllParams
}

type GetAllQueryHandler struct {
	UserRepository userRepository.Repository
}

func (h *GetAllQueryHandler) Execute(query GetAllQuery) ([]*userModels.Model, int, error) {
	users, count, err := h.UserRepository.GetAll(query.Params)

	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}
