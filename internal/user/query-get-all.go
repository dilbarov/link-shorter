package user

import "context"

type GetAllQuery struct {
	Params *GetAllParams
}

type GetAllQueryHandler struct {
	UserRepository Repository
}

func NewGetAllQueryHandler(repo Repository) *GetAllQueryHandler {
	return &GetAllQueryHandler{
		UserRepository: repo,
	}
}

func (h *GetAllQueryHandler) Handle(ctx context.Context, query GetAllQuery) ([]*Model, int, error) {
	users, count, err := h.UserRepository.GetAll(query.Params)

	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}
