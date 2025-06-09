package click

import (
	"context"
	clickPayloads "link-shorter/internal/click/payloads"
	clickRepository "link-shorter/internal/click/repository"
	clickResponses "link-shorter/internal/click/responses"
	"link-shorter/pkg/date"
)

const GetStatsQueryName = "click.get-stats"

type GetStatsQuery struct {
	Params clickPayloads.GetStatsParams
}

func (q GetStatsQuery) Name() string {
	return GetStatsQueryName
}

type GetStatsQueryHandler struct {
	ClickRepository clickRepository.Repository
}

func NewGetStatsQueryHandler(clickRepository clickRepository.Repository) *GetStatsQueryHandler {
	return &GetStatsQueryHandler{
		ClickRepository: clickRepository,
	}
}

func (h *GetStatsQueryHandler) Handle(ctx context.Context, query GetStatsQuery) ([]*clickResponses.CountStatItem, error) {
	startDate := date.ToStartOfDay(query.Params.StartDate)
	endDate := date.ToEndOfDay(query.Params.EndDate)

	stats, err := h.ClickRepository.GetStatsByLink(query.Params.LinkId, &startDate, &endDate, query.Params.Timezone)
	if err != nil {
		return nil, err
	}

	return stats, nil
}
