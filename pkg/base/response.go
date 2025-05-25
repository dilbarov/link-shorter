package base

import "github.com/google/uuid"

type PublicResponse struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt *string   `json:"deletedAt"`
}
