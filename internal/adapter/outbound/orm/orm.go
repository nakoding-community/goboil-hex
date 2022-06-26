package orm

import (
	"github.com/nakoding-community/golang-boilerplate/internal/model"

	"gorm.io/gorm"
)

type (
	Orm struct {
		User        model.User
		Transaction model.Transaction
	}
)

func New(connection *gorm.DB) *Orm {
	return &Orm{
		User:        NewUser(connection),
		Transaction: NewTransaction(connection),
	}
}
