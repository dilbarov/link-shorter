package click

import (
	"time"
)

type Repository interface {
	GetCountByLink(linkId string, startDate *time.Time, endDate *time.Time) (int64, error)
	GetStatsByLink(linkId string, startDate *time.Time, endDate *time.Time, tz *string) ([]*CountStatItem, error)
	Create(click *Model) (*Model, error)
}
