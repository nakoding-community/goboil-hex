package usecase

import (
	"github.com/nakoding-community/golang-boilerplate/entity"
	"github.com/nakoding-community/golang-boilerplate/internal/model"
)

type (
	User interface {
		Base[entity.User]
	}

	user struct {
		Base[entity.User]

		modelUser model.User
	}
)

func NewUser(
	modelUser model.User,
) User {
	return &user{
		Base:      NewBase[entity.User](modelUser),
		modelUser: modelUser,
	}
}

func (u *user) Custom() (string, error) {
	return "", nil
}
