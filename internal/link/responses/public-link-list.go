package link

import (
	linkModels "link-shorter/internal/link/models"
	"link-shorter/pkg/res"
)

type PublicLinkList = res.Paginated[*PublicLink]

func NewPublicLinkList(links []*linkModels.Model, count *int) *PublicLinkList {
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
