package queries

import (
	linkModels "link-shorter/internal/link/models"
	linkPayloads "link-shorter/internal/link/payloads"
	linkRepository "link-shorter/internal/link/repository"
)

type GetByIdQuery struct {
	Params *linkPayloads.GetByIDParams
}

type GetByIdQueryHandler struct {
	LinkRepository linkRepository.Repository
}

func (h *GetByIdQueryHandler) Execute(query GetByIdQuery) (*linkModels.Model, error) {
	link, err := h.LinkRepository.GetById(query.Params.Id)

	if err != nil {
		return nil, err
	}

	return link, nil
}
