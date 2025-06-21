package link

import "context"

const GetByHashQueryName = "link.get-by-hash"

type GetByHashQuery struct {
	Params *GetByHashParams
}

func (q GetByHashQuery) Name() string {
	return GetByHashQueryName
}

type GetByHashQueryHandler struct {
	LinkRepository Repository
}

func NewGetByHashQueryHandler(linkRepository Repository) *GetByHashQueryHandler {
	return &GetByHashQueryHandler{
		LinkRepository: linkRepository,
	}
}

func (h *GetByHashQueryHandler) Handle(ctx context.Context, query GetByHashQuery) (*Model, error) {
	link, err := h.LinkRepository.GetByHash(query.Params.Hash)

	if err != nil {
		return nil, err
	}

	return link, nil
}
