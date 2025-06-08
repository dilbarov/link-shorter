package link

import (
	"context"
	linkModels "link-shorter/internal/link/models"
	linkPayloads "link-shorter/internal/link/payloads"
	linkRepository "link-shorter/internal/link/repository"
)

const GetByHIdQueryName = "link.get-by-id"

type GetByIdQuery struct {
	Params *linkPayloads.GetByIDParams
}

func (q GetByIdQuery) Name() string {
	return GetByHIdQueryName
}

type GetByIdQueryHandler struct {
	LinkRepository linkRepository.Repository
}

func NewGetByIdQueryHandler(linkRepository linkRepository.Repository) *GetByIdQueryHandler {
	return &GetByIdQueryHandler{
		LinkRepository: linkRepository,
	}
}

func (h *GetByIdQueryHandler) Handle(ctx context.Context, query GetByIdQuery) (*linkModels.Model, error) {
	link, err := h.LinkRepository.GetById(query.Params.Id)

	if err != nil {
		return nil, err
	}

	return link, nil
}
