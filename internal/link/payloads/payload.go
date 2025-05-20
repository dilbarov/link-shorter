package payloads

type CreateLinkRequest struct {
	Url string `json:"url"`
}

type UpdateLinkRequest struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}
