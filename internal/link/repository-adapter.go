package link

import "link-shorter/pkg/db"

type PostgresLinkRepository struct {
	Database *db.Db
}

func NewPostgresLinkRepository(database *db.Db) Repository {
	return &PostgresLinkRepository{
		Database: database,
	}
}

func (repo *PostgresLinkRepository) GetByHash(hash string) (*Model, error) {
	var link Model
	result := repo.Database.DB.First(&link, "hash = ?", hash)

	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}

func (repo *PostgresLinkRepository) GetById(id string) (*Model, error) {
	var link Model
	result := repo.Database.DB.First(&link, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}

func (repo *PostgresLinkRepository) GetAll(query *GetAllParams) ([]*Model, int, error) {
	dbQuery := repo.Database.Model(&Model{})

	if query.Search != nil && *query.Search != "" {
		searchPattern := "%" + *query.Search + "%"
		dbQuery = dbQuery.Where("url ILIKE ? OR hash ILIKE ? and deleted_at IS NULL", searchPattern, searchPattern)
	}

	items, total, err := db.FindAndCount[Model](dbQuery, db.CommonFilter{
		Limit:  *query.Limit,
		Offset: *query.Offset,
	})

	return items, int(total), err
}

func (repo *PostgresLinkRepository) Create(link *Model) (*Model, error) {
	result := repo.Database.Create(link)

	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}

func (repo *PostgresLinkRepository) Update(link *Model) (*Model, error) {
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
