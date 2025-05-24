package link

import (
	linkModels "link-shorter/internal/link/models"
	linkPayloads "link-shorter/internal/link/payloads"
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

func (repo *PostgresLinkRepository) GetAll(query *linkPayloads.GetAllParams) ([]*linkModels.Model, int, error) {
	var (
		links []*linkModels.Model
		total int64
	)

	dbQuery := repo.Database.Model(&linkModels.Model{})

	if query.Search != nil && *query.Search != "" {
		searchPattern := "%" + *query.Search + "%"
		dbQuery = dbQuery.Where("url ILIKE ? OR hash ILIKE ?", searchPattern, searchPattern)
	}

	if err := dbQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if *query.Limit > 0 {
		dbQuery = dbQuery.Limit(*query.Limit)
	}

	if *query.Offset > 0 {
		dbQuery = dbQuery.Offset(*query.Offset)
	}

	if err := dbQuery.Find(&links).Error; err != nil {
		return nil, 0, err
	}

	return links, int(total), nil
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
