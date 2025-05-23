package link

type GetByHashParams struct {
	Hash string
}

type GetByIDParams struct {
	ID string
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
