package user

import "context"

type GetByEmailQuery struct {
	Params *GetByEmailParams
}

type GetByEmailQueryHandler struct {
	UserRepository Repository
}

func NewGetByEmailQueryHandler(repo Repository) *GetByEmailQueryHandler {
	return &GetByEmailQueryHandler{
		UserRepository: repo,
	}
}

func (h GetByEmailQueryHandler) Handle(ctx context.Context, query GetByEmailQuery) (*Model, error) {
	user, err := h.UserRepository.GetByEmail(query.Params.Email)

	if err != nil {
		return nil, err
	}

	return user, nil
}
