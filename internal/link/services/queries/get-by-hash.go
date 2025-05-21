package queries

import (
	"link-shorter/internal/link/models"
	"link-shorter/internal/link/payloads"
	"link-shorter/internal/link/repository"
)

type GetByHashQuery struct {
	Params *payloads.LinkGetByHashParams
}

type GetByHashQueryHandler struct {
	LinkRepository repository.LinkRepository
}

func (h *GetByHashQueryHandler) Execute(query GetByHashQuery) (*models.LinkModel, error) {
	link, err := h.LinkRepository.GetByHash(query.Params.Hash)

	if err != nil {
		return nil, err
	}

	return link, nil
}
