package link

import (
	linkModels "link-shorter/internal/link/models"
	"link-shorter/pkg/base"
	"time"
)

type PublicLink struct {
	base.PublicResponse
	Url  string `json:"url"`
	Hash string `json:"hash"`
}

func NewPublicLink(m *linkModels.Model) *PublicLink {
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
