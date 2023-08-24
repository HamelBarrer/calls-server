package user

import (
	"github.com/HamelBarrer/calls-server/internal/middlewares"
	"github.com/labstack/echo/v4"
)

type Routers interface {
	SetupConfig(Repository)
}

type Router struct {
	e *echo.Echo
}

func NewRouter(e *echo.Echo) Routers {
	return &Router{e}
}

func (ro *Router) SetupConfig(r Repository) {
	c := NewController(r)

	ro.e.GET("/api/v1/users/:user_id", c.GetUser, middlewares.ValidAuth)
	ro.e.GET("/api/v1/users", c.GetUsers, middlewares.ValidAuth)
	ro.e.POST("/api/v1/users", c.CreateUser)
	ro.e.PUT("/api/v1/users/:user_id", c.UpdatedUser, middlewares.ValidAuth)
}
