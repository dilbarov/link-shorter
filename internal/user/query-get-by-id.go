package user

import (
	"context"
)

type GetByIdQuery struct {
	Params *GetByIdParams
}

type GetByIdQueryHandler struct {
	UserRepository Repository
}

func NewGetByIdQueryHandler(repo Repository) *GetByIdQueryHandler {
	return &GetByIdQueryHandler{
		UserRepository: repo,
	}
}

func (h *GetByIdQueryHandler) Handle(ctx context.Context, query GetByIdQuery) (*Model, error) {
	user, err := h.UserRepository.GetById(query.Params.Id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
