package response

import (
	"github.com/nakoding-community/golang-boilerplate/dto"
)

type Meta struct {
	Success bool                `json:"success" default:"true"`
	Message string              `json:"message" default:"true"`
	Info    *dto.PaginationInfo `json:"info"`
}

type responseHelper struct {
	Error   errorHelper
	Success successHelper
}

var Constant responseHelper = responseHelper{
	Error:   errorConstant,
	Success: successConstant,
}
