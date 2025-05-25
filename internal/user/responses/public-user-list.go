package user

import (
	userModels "link-shorter/internal/user/models"
	"link-shorter/pkg/res"
)

type PublicUserList = res.Paginated[*PublicUser]

func NewPublicUserList(users []*userModels.Model, count *int) *PublicUserList {
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
