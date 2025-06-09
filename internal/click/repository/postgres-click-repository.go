package click

import (
	"errors"
	"fmt"
	clickModels "link-shorter/internal/click/models"
	clickResponses "link-shorter/internal/click/responses"
	"link-shorter/pkg/db"
	"time"
)

type PostgresClickRepository struct {
	Database *db.Db
}

func NewPostgresClickRepository(database *db.Db) Repository {
	return &PostgresClickRepository{
		Database: database,
	}
}

func (repo *PostgresClickRepository) Create(click *clickModels.Model) (*clickModels.Model, error) {
	result := repo.Database.Create(click)

	if result.Error != nil {
		return nil, result.Error
	}

	return click, nil
}

func (repo *PostgresClickRepository) GetCountByLink(linkId string, startDate *time.Time, endDate *time.Time) (int64, error) {
	var count int64 = 0

	dbQuery := repo.Database.Model(&clickModels.Model{})

	dbQuery.
		Select("DATE(created_at) as date, COUNT(*) as count").
		Where("link_id = ? AND deleted_at IS NULL", linkId)

	if startDate != nil && endDate != nil {
		if endDate.Before(*startDate) {
			return 0, errors.New("endDate cannot be before startDate")
		}
	}

	if startDate != nil {
		dbQuery.Where("created_at >= ?", startDate)
	}

	if endDate != nil {
		dbQuery.Where("created_at <= ?", endDate)
	}

	//dbQuery.Order("created_at DESC").Group("created_at").Count(&count)
	dbQuery.Count(&count)

	return count, nil
}

func (repo *PostgresClickRepository) GetStatsByLink(linkId string, startDate *time.Time, endDate *time.Time, tz *string) ([]*clickResponses.CountStatItem, error) {
	defaultTZ := "America/New_York"

	if tz == nil {
		tz = &defaultTZ
	}

	if startDate != nil && endDate != nil {
		if endDate.Before(*startDate) {
			return nil, errors.New("endDate cannot be before startDate")
		}
	}

	var stats []*clickResponses.CountStatItem

	dbQuery := repo.Database.Model(&clickModels.Model{})

	dbQuery.
		Select(fmt.Sprintf("DATE(created_at AT TIME ZONE 'UTC' AT TIME ZONE '%s') as Date, COUNT(*) as Count", tz)).
		Where("link_id = ? AND deleted_at IS NULL", linkId)

	if startDate != nil {
		dbQuery.Where("created_at >= ?", startDate)
	}

	if endDate != nil {
		dbQuery.Where("created_at <= ?", endDate)
	}

	err := dbQuery.
		Group(fmt.Sprintf("DATE(created_at AT TIME ZONE 'UTC' AT TIME ZONE '%s')", tz)).
		Order("Date").
		Scan(&stats).Error

	if err != nil {
		return nil, err
	}

	return stats, nil
}
