package repository

import (
	"link-shorter/internal/link/models"
)

type LinkRepository interface {
	GetByHash(hash string) (*models.LinkModel, error)
	GetById(id string) (*models.LinkModel, error)
	Create(link *models.LinkModel) (*models.LinkModel, error)
	Update(id string, link *models.LinkModel) (*models.LinkModel, error)
}
