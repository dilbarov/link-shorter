package repository

import (
	"link-shorter/internal/link/models"
	"link-shorter/pkg/db"
)

type PostgresLinkRepository struct {
	Database *db.Db
}

func NewPostgresLinkRepository(database *db.Db) LinkRepository {
	return &PostgresLinkRepository{
		Database: database,
	}
}

func (repo *PostgresLinkRepository) Create(link *models.LinkModel) (*models.LinkModel, error) {
	return nil, nil
}

func (repo *PostgresLinkRepository) Update(id string, link *models.LinkModel) (*models.LinkModel, error) {
	return nil, nil
}
