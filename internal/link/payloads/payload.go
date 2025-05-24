package link

type GetByHashParams struct {
	Hash string
}

type GetByIDParams struct {
	Id string
}

type GetAllParams struct {
	Limit  *int    `schema:"limit,default:10"`
	Offset *int    `schema:"offset,default:0"`
	Search *string `schema:"search"`
}

type CreatePayload struct {
	Url string `json:"url" validate:"required,url"`
}

type UpdatePayload struct {
	Id   string
	Url  *string
	Hash *string
}

type UpdateRequest struct {
	Url  *string `json:"url" validate:"omitempty,url"`
	Hash *string `json:"hash" validate:"omitempty"`
}
