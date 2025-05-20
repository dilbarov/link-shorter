package link

import (
	"link-shorter/internal/link/models"
	"link-shorter/internal/link/payloads"
	"link-shorter/internal/link/repository"
)

type CreateCommand struct {
	Payload *payloads.CreateLinkRequest
}

type CreateHandler struct {
	Repo repository.LinkRepository
}

func (h *CreateHandler) Execute(cmd CreateCommand) (*models.LinkModel, error) {
	link := models.NewLink(cmd.Payload.Url)
	if _, err := h.Repo.Create(link); err != nil {
		return nil, err
	}
	return link, nil
}
