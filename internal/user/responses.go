package user

import (
	"link-shorter/pkg/base"
	"link-shorter/pkg/res"
	"time"
)

type PublicUser struct {
	base.PublicResponse
	Name  *string `json:"name" validate:"omitempty,min=1,max=32"`
	Email string  `json:"email" validate:"required,email"`
}

func NewPublicUser(m *Model) *PublicUser {
	var deletedAt *string
	if m.DeletedAt.Valid {
		s := m.DeletedAt.Time.Format(time.RFC3339)
		deletedAt = &s
	}

	return &PublicUser{
		PublicResponse: base.PublicResponse{
			Id:        m.Id,
			CreatedAt: m.CreatedAt.Format(time.RFC3339),
			UpdatedAt: m.UpdatedAt.Format(time.RFC3339),
			DeletedAt: deletedAt,
		},
		Name:  m.Name,
		Email: m.Email,
	}
}

type PublicUserList = res.Paginated[*PublicUser]

func NewPublicUserList(users []*Model, count *int) *PublicUserList {
	var items []*PublicUser

	for _, user := range users {
		items = append(items, NewPublicUser(user))
	}

	if items == nil {
		items = []*PublicUser{}
	}

	return &PublicUserList{
		Items: items,
		Count: *count,
	}
}
