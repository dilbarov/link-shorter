package link

import (
	linkModels "link-shorter/internal/link/models"
	"link-shorter/pkg/db"
)

type PostgresLinkRepository struct {
	Database *db.Db
}

func NewPostgresLinkRepository(database *db.Db) Repository {
	return &PostgresLinkRepository{
		Database: database,
	}
}

func (repo *PostgresLinkRepository) GetByHash(hash string) (*linkModels.Model, error) {
	var link linkModels.Model
	result := repo.Database.DB.First(&link, "hash = ?", hash)

	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}

func (repo *PostgresLinkRepository) GetById(id string) (*linkModels.Model, error) {
	var link linkModels.Model
	result := repo.Database.DB.First(&link, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}

func (repo *PostgresLinkRepository) Create(link *linkModels.Model) (*linkModels.Model, error) {
	result := repo.Database.Create(link)

	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}

func (repo *PostgresLinkRepository) Update(link *linkModels.Model) (*linkModels.Model, error) {
	result := repo.Database.DB.Updates(link)

	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}

func (repo *PostgresLinkRepository) Delete(id string) error {
	link, err := repo.GetById(id)

	if err != nil {
		return err
	}

	repo.Database.DB.Delete(link)
	return nil
}
