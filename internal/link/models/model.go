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
	return &Model{
		Url:  url,
		Hash: RandomStringRunes(10),
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwyzABCDEFGHIJKLMOPQRSTUVWXYZ")

func RandomStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
