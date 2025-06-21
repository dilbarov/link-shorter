package click

import (
	"context"
	"link-shorter/pkg/date"
)

type GetStatsQuery struct {
	Params GetStatsParams
}

type GetStatsQueryHandler struct {
	ClickRepository Repository
}

func NewGetStatsQueryHandler(clickRepository Repository) *GetStatsQueryHandler {
	return &GetStatsQueryHandler{
		ClickRepository: clickRepository,
	}
}

func (h *GetStatsQueryHandler) Handle(ctx context.Context, query GetStatsQuery) ([]*CountStatItem, error) {
	startDate := date.ToStartOfDay(query.Params.StartDate)
	endDate := date.ToEndOfDay(query.Params.EndDate)

	stats, err := h.ClickRepository.GetStatsByLink(query.Params.LinkId, &startDate, &endDate, query.Params.Timezone)
	if err != nil {
		return nil, err
	}

	return stats, nil
}
