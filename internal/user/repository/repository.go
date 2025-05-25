package user

import (
	userModels "link-shorter/internal/user/models"
	userPayloads "link-shorter/internal/user/payloads"
)

type Repository interface {
	GetById(id string) (*userModels.Model, error)
	GetByEmail(email string) (*userModels.Model, error)
	GetAll(query *userPayloads.GetAllParams) ([]*userModels.Model, int, error)
	Create(model *userModels.Model) (*userModels.Model, error)
	Update(model *userModels.Model) (*userModels.Model, error)
	Delete(id string) error
}
