package link

import (
	linkModels "link-shorter/internal/link/models"
	linkPayloads "link-shorter/internal/link/payloads"
	linkRepository "link-shorter/internal/link/repository"
)

type GetAllQuery struct {
	Params *linkPayloads.GetAllParams
}

type GetAllQueryHandler struct {
	LinkRepository linkRepository.Repository
}

func (h *GetAllQueryHandler) Execute(query GetAllQuery) ([]*linkModels.Model, int, error) {
	links, count, err := h.LinkRepository.GetAll(query.Params)

	if err != nil {
		return nil, 0, err
	}

	return links, count, err
}
