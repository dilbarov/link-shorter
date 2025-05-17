package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"link-shorter/configs"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *configs.DbConfig) *Db {
	db, err := gorm.Open(postgres.Open(conf.Dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &Db{db}
}
