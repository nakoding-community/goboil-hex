package usecase

import "github.com/nakoding-community/golang-boilerplate/internal/model"

type (
	Usecase struct {
		UsecaseUser User
	}
)

func New(
	modelUser model.User,
) *Usecase {
	return &Usecase{
		UsecaseUser: NewUser(modelUser),
	}
}
