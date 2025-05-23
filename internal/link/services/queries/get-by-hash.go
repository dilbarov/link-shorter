package link

import (
	linkModels "link-shorter/internal/link/models"
	linkPayloads "link-shorter/internal/link/payloads"
	linkRepository "link-shorter/internal/link/repository"
)

type GetByHashQuery struct {
	Params *linkPayloads.GetByHashParams
}

type GetByHashQueryHandler struct {
	LinkRepository linkRepository.Repository
}

func (h *GetByHashQueryHandler) Execute(query GetByHashQuery) (*linkModels.Model, error) {
	link, err := h.LinkRepository.GetByHash(query.Params.Hash)

	if err != nil {
		return nil, err
	}

	return link, nil
}
