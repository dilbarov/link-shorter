package models

import (
	"gorm.io/gorm"
	"math/rand"
)

type LinkModel struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

func NewLink(url string) *LinkModel {
	return &LinkModel{
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
