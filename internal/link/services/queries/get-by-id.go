package queries

import (
	"link-shorter/internal/link/models"
	"link-shorter/internal/link/payloads"
	"link-shorter/internal/link/repository"
)

type GetByIdQuery struct {
	Params *payloads.LinkGetByIDParams
}

type GetByIdQueryHandler struct {
	LinkRepository repository.LinkRepository
}

func (h *GetByIdQueryHandler) Execute(query GetByIdQuery) (*models.LinkModel, error) {
	link, err := h.LinkRepository.GetById(query.Params.ID)

	if err != nil {
		return nil, err
	}

	return link, nil
}
