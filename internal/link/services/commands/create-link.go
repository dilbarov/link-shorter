package link

import (
	"link-shorter/internal/link/models"
	"link-shorter/internal/link/payloads"
	"link-shorter/internal/link/repository"
)

type CreateCommand struct {
	Payload *payloads.LinkCreateRequest
}

type CreateHandler struct {
	LinkRepository repository.LinkRepository
}

func (h *CreateHandler) Execute(cmd CreateCommand) (*models.LinkModel, error) {
	link := models.NewLink(cmd.Payload.Url)
	createdLink, err := h.LinkRepository.Create(link)
	if err != nil {
		return nil, err
	}
	return createdLink, nil
}
