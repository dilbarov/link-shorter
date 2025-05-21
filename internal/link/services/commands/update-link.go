package link

import (
	"link-shorter/internal/link/models"
	"link-shorter/internal/link/payloads"
	"link-shorter/internal/link/repository"
)

type UpdateCommand struct {
	Payload *payloads.LinkUpdateRequest
}

type UpdateCommandHandler struct {
	LinkRepository repository.LinkRepository
}

func (h *UpdateCommandHandler) Execute(cmd UpdateCommand) (*models.LinkModel, error) {
	link := models.NewLink(cmd.Payload.Url)
	result, err := h.LinkRepository.Update(cmd.Payload.Id, link)
	if err != nil {
		return nil, err
	}
	return result, nil
}
