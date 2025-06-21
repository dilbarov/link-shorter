package link

type Repository interface {
	GetByHash(hash string) (*Model, error)
	GetById(id string) (*Model, error)
	GetAll(query *GetAllParams) ([]*Model, int, error)
	Create(link *Model) (*Model, error)
	Update(link *Model) (*Model, error)
	Delete(id string) error
}
