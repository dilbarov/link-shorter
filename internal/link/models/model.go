package link

import (
	"link-shorter/pkg/base"
	"math/rand"
)

type Model struct {
	base.Model
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

func NewLink(url string) *Model {
	link := &Model{Url: url}
	link.GenerateHash()
	return link
}

func (link *Model) GenerateHash() {
	link.Hash = RandomStringRunes(10)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwyzABCDEFGHIJKLMOPQRSTUVWXYZ")

func RandomStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
