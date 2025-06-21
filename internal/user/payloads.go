package user

type GetByIdParams struct {
	Id string
}

type GetByEmailParams struct {
	Email string
}

type GetAllParams struct {
	Limit  *int    `schema:"limit,default:10"`
	Offset *int    `schema:"offset,default:0"`
	Search *string `schema:"search"`
}

type CreatePayload struct {
	Name     string `validate:"required,min=1,max=50"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type UpdatePayload struct {
	Id   string
	Name *string
}

type UpdateRequest struct {
	Name *string `validate:"omitempty,min=1,max=50"`
}
