package model

import (
	"context"

	"github.com/nakoding-community/golang-boilerplate/entity"
)

type (
	User interface {
		Base[entity.User]
		Some(ctx context.Context, id int) error
	}
)
