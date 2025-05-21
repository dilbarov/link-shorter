package payloads

type LinkCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type LinkUpdateRequest struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}
