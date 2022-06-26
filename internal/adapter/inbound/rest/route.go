package rest

import (
	"fmt"
	"net/http"

	"github.com/nakoding-community/golang-boilerplate/internal/adapter/inbound/rest/handler/user"

	"github.com/labstack/echo/v4"
)

func (r *rest) InitRoute() {
	r.e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", "github.com/nakoding-community/golang-boilerplate", "0.0.1")
		return c.String(http.StatusOK, message)
	})

	r.initV1Route(r.e.Group("/v1"))
}

func (r *rest) initV1Route(g *echo.Group) {
	user.NewHandler(
		r.usecaseUser,
	).Route(g.Group("/users"))
}
