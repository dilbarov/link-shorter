package click

import (
	"context"
	clickPayloads "link-shorter/internal/click/payloads"
	clickRepository "link-shorter/internal/click/repository"
)

const GetCountQueryName = "click.get-count"

type GetCountQuery struct {
	Params *clickPayloads.GetCountParams
}

func (q GetCountQuery) Name() string {
	return GetCountQueryName
}

type GetCountQueryHandler struct {
	ClickRepository clickRepository.Repository
}

func NewGetCountQueryHandler(clickRepository clickRepository.Repository) *GetCountQueryHandler {
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
