package click

import (
	"context"
)

type GetCountQuery struct {
	Params *GetCountParams
}

type GetCountQueryHandler struct {
	ClickRepository Repository
}

func NewGetCountQueryHandler(clickRepository Repository) *GetCountQueryHandler {
	return &GetCountQueryHandler{
		ClickRepository: clickRepository,
	}
}

func (h *GetCountQueryHandler) Handle(ctx context.Context, query GetCountQuery) (int64, error) {
	count, err := h.ClickRepository.GetCountByLink(query.Params.LinkId, nil, nil)
	if err != nil {
		return 0, err
	}

	return count, nil
}
