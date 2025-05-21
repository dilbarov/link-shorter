package link

import (
	"link-shorter/internal/link/models"
	"link-shorter/internal/link/payloads"
	"link-shorter/internal/link/repository"
)

type UpdateCommand struct {
	Payload *payloads.LinkUpdateRequest
}

type UpdateHandler struct {
	Repo repository.LinkRepository
}

func (h *UpdateHandler) Execute(cmd UpdateCommand) (*models.LinkModel, error) {
	link := models.NewLink(cmd.Payload.Url)
	result, err := h.Repo.Update(cmd.Payload.Id, link)
	if err != nil {
		return nil, err
	}
	return result, nil
}
