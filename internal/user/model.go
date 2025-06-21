package user

import "link-shorter/pkg/base"

type Model struct {
	base.Model
	Name         *string
	Email        string `gorm:"unique,index"`
	PasswordHash *string
}

func (*Model) TableName() string {
	return "users"
}

func NewUser(email string, password *string, name *string) *Model {
	user := &Model{
		Email: email,
	}

	if password != nil {
		user.GenerateHash(password)
	}

	if name != nil {
		user.Name = name
	}

	return user
}

func (user *Model) GenerateHash(password *string) {

}
