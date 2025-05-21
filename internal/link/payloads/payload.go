package link

type GetByHashParams struct {
	Hash string
}

type GetByIDParams struct {
	ID string
}

type CreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type UpdateRequest struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}
