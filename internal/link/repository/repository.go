package link

import (
	linkModels "link-shorter/internal/link/models"
)

type Repository interface {
	GetByHash(hash string) (*linkModels.Model, error)
	GetById(id string) (*linkModels.Model, error)
	Create(link *linkModels.Model) (*linkModels.Model, error)
	Update(link *linkModels.Model) (*linkModels.Model, error)
	Delete(id string) error
}
