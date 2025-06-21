package link

import "context"

const GetAllLinksQueryName = "link.get-all"

type GetAllQuery struct {
	Params *GetAllParams
}

func (q GetAllQuery) Name() string {
	return GetAllLinksQueryName
}

type GetAllQueryHandler struct {
	LinkRepository Repository
}

func NewGetAllQueryHandler(linkRepository Repository) *GetAllQueryHandler {
	return &GetAllQueryHandler{
		LinkRepository: linkRepository,
	}
}

func (h *GetAllQueryHandler) Handle(ctx context.Context, query GetAllQuery) ([]*Model, int, error) {
	links, count, err := h.LinkRepository.GetAll(query.Params)

	if err != nil {
		return nil, 0, err
	}

	return links, count, err
}
