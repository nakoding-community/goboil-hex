package adapter

import (
	"github.com/nakoding-community/golang-boilerplate/internal/adapter/inbound/rest"
	"github.com/nakoding-community/golang-boilerplate/internal/usecase"
)

type (
	InBound struct {
		Rest rest.Rest
	}
)

func NewInBound(
	usecaseUser usecase.User,
) *InBound {
	restInstance := rest.New(usecaseUser)

	return &InBound{
		Rest: restInstance,
	}
}
