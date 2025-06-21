package link

import (
	"link-shorter/pkg/base"
	"link-shorter/pkg/res"
	"time"
)

type PublicLink struct {
	base.PublicResponse
	Url  string `json:"url"`
	Hash string `json:"hash"`
}

func NewPublicLink(m *Model) *PublicLink {
	var deletedAt *string
	if m.DeletedAt.Valid {
		s := m.DeletedAt.Time.Format(time.RFC3339)
		deletedAt = &s
	}

	return &PublicLink{
		PublicResponse: base.PublicResponse{
			Id:        m.Id,
			CreatedAt: m.CreatedAt.Format(time.RFC3339),
			UpdatedAt: m.UpdatedAt.Format(time.RFC3339),
			DeletedAt: deletedAt,
		},
		Url:  m.Url,
		Hash: m.Hash,
	}
}

type PublicLinkList = res.Paginated[*PublicLink]

func NewPublicLinkList(links []*Model, count *int) *PublicLinkList {
	var items []*PublicLink

	for _, link := range links {
		items = append(items, NewPublicLink(link))
	}

	if items == nil {
		items = []*PublicLink{}
	}

	return &PublicLinkList{
		Items: items,
		Count: *count,
	}
}
