package repository

import (
	"link-shorter/internal/link/models"
)

type LinkRepository interface {
	Create(link *models.LinkModel) (*models.LinkModel, error)
	Update(id string, link *models.LinkModel) (*models.LinkModel, error)
}
