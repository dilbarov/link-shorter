package user

import (
	userModels "link-shorter/internal/user/models"
	userPayloads "link-shorter/internal/user/payloads"
	userRepository "link-shorter/internal/user/repository"
)

type GetByEmailQuery struct {
	Params *userPayloads.GetByEmailParams
}

type GetByEmailQueryHandler struct {
	UserRepository userRepository.Repository
}

func (h GetByEmailQueryHandler) Execute(query GetByEmailQuery) (*userModels.Model, error) {
	user, err := h.UserRepository.GetByEmail(query.Params.Email)

	if err != nil {
		return nil, err
	}

	return user, nil
}
