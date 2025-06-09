package click

import (
	clickModels "link-shorter/internal/click/models"
	clickResponses "link-shorter/internal/click/responses"
	"time"
)

type Repository interface {
	GetCountByLink(linkId string, startDate *time.Time, endDate *time.Time) (int64, error)
	GetStatsByLink(linkId string, startDate *time.Time, endDate *time.Time, tz *string) ([]*clickResponses.CountStatItem, error)
	Create(click *clickModels.Model) (*clickModels.Model, error)
}
