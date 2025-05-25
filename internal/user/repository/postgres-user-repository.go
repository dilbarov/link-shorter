package user

import (
	userModels "link-shorter/internal/user/models"
	userPayloads "link-shorter/internal/user/payloads"
	"link-shorter/pkg/db"
)

type PostgresUserRepository struct {
	Database *db.Db
}

func NewPostgresUserRepository(database *db.Db) *PostgresUserRepository {
	return &PostgresUserRepository{
		Database: database,
	}
}

func (repo *PostgresUserRepository) GetById(id string) (*userModels.Model, error) {
	var user userModels.Model
	result := repo.Database.DB.First(&user, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo *PostgresUserRepository) GetByEmail(email string) (*userModels.Model, error) {
	var user userModels.Model
	result := repo.Database.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo *PostgresUserRepository) GetAll(query *userPayloads.GetAllParams) ([]*userModels.Model, int, error) {
	var (
		users []*userModels.Model
		total int64
	)

	dbQuery := repo.Database.Model(&userModels.Model{})

	if query.Search != nil && *query.Search != "" {
		searchPattern := "%" + *query.Search + "%"
		dbQuery = dbQuery.Where("email LIKE ?", searchPattern)
	}

	if err := dbQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if *query.Limit > 0 {
		dbQuery.Limit(*query.Limit)
	}

	if *query.Offset > 0 {
		dbQuery.Offset(*query.Offset)
	}

	if err := dbQuery.Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, int(total), nil
}

func (repo *PostgresUserRepository) Create(user *userModels.Model) (*userModels.Model, error) {
	if err := repo.Database.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *PostgresUserRepository) Update(user *userModels.Model) (*userModels.Model, error) {
	if err := repo.Database.DB.Updates(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *PostgresUserRepository) Delete(id string) error {
	link, err := repo.GetById(id)

	if err != nil {
		return err
	}

	repo.Database.DB.Delete(link)
	return nil
}
