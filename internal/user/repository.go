package user

type Repository interface {
	GetById(id string) (*Model, error)
	GetByEmail(email string) (*Model, error)
	GetAll(query *GetAllParams) ([]*Model, int, error)
	Create(model *Model) (*Model, error)
	Update(model *Model) (*Model, error)
	Delete(id string) error
}
