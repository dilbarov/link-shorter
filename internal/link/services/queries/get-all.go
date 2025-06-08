package link

import (
	"context"
	linkModels "link-shorter/internal/link/models"
	linkPayloads "link-shorter/internal/link/payloads"
	linkRepository "link-shorter/internal/link/repository"
	"link-shorter/pkg/db"
)

const GetAllLinksQueryName = "link.get-all"

type GetAllQuery struct {
	Params *linkPayloads.GetAllParams
}

func (q GetAllQuery) Name() string {
	return GetAllLinksQueryName
}

type GetAllQueryHandler struct {
	LinkRepository linkRepository.Repository
}

func NewGetAllQueryHandler(linkRepository linkRepository.Repository) *GetAllQueryHandler {
	return &GetAllQueryHandler{
		LinkRepository: linkRepository,
	}
}

func (h *GetAllQueryHandler) Handle(ctx context.Context, query GetAllQuery) (*db.ListResult[*linkModels.Model], error) {
	links, count, err := h.LinkRepository.GetAll(query.Params)

	if err != nil {
		return nil, err
	}

	result := db.NewListResult(links, count)

	return result, err
}
