package customauth

import "github.com/labstack/echo/v4"

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

	ro.e.POST("/api/v1/auth", c.Login)
}
