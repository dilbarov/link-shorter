package link

import (
	"github.com/google/uuid"
	linkModels "link-shorter/internal/link/models"
	"time"
)

type PublicLink struct {
	Id        uuid.UUID `json:"id"`
	Url       string    `json:"url"`
	Hash      string    `json:"hash"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt *string   `json:"deletedAt"`
}

func NewPublicLink(m *linkModels.Model) *PublicLink {
	var deletedAt *string
	if m.DeletedAt.Valid {
		s := m.DeletedAt.Time.Format(time.RFC3339)
		deletedAt = &s
	}

	return &PublicLink{
		Id:        m.Id,
		Url:       m.Url,
		Hash:      m.Hash,
		CreatedAt: m.CreatedAt.Format(time.RFC3339),
		UpdatedAt: m.UpdatedAt.Format(time.RFC3339),
		DeletedAt: deletedAt,
	}
}
