package link

import "context"

const GetByHIdQueryName = "link.get-by-id"

type GetByIdQuery struct {
	Params *GetByIDParams
}

func (q GetByIdQuery) Name() string {
	return GetByHIdQueryName
}

type GetByIdQueryHandler struct {
	LinkRepository Repository
}

func NewGetByIdQueryHandler(linkRepository Repository) *GetByIdQueryHandler {
	return &GetByIdQueryHandler{
		LinkRepository: linkRepository,
	}
}

func (h *GetByIdQueryHandler) Handle(ctx context.Context, query GetByIdQuery) (*Model, error) {
	link, err := h.LinkRepository.GetById(query.Params.Id)

	if err != nil {
		return nil, err
	}

	return link, nil
}
