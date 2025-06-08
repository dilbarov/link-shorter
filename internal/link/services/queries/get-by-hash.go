package link

import (
	"context"
	linkModels "link-shorter/internal/link/models"
	linkPayloads "link-shorter/internal/link/payloads"
	linkRepository "link-shorter/internal/link/repository"
)

const GetByHashQueryName = "link.get-by-hash"

type GetByHashQuery struct {
	Params *linkPayloads.GetByHashParams
}

func (q GetByHashQuery) Name() string {
	return GetByHashQueryName
}

type GetByHashQueryHandler struct {
	LinkRepository linkRepository.Repository
}

func NewGetByHashQueryHandler(linkRepository linkRepository.Repository) *GetByHashQueryHandler {
	return &GetByHashQueryHandler{
		LinkRepository: linkRepository,
	}
}

func (h *GetByHashQueryHandler) Handle(ctx context.Context, query GetByHashQuery) (*linkModels.Model, error) {
	link, err := h.LinkRepository.GetByHash(query.Params.Hash)

	if err != nil {
		return nil, err
	}

	return link, nil
}
