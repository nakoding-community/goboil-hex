package model

import (
	"context"

	"github.com/nakoding-community/golang-boilerplate/entity"
)

type (
	Transaction interface {
		Base[entity.Transaction]
		Some(ctx context.Context, id int) error
	}
)
