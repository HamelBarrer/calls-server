package commentary

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

	ro.e.GET("/api/v1/commentaries/:commentary_id", c.GetCommentary, middlewares.ValidAuth)
	ro.e.GET("/api/v1/commentaries", c.GetCommentaries, middlewares.ValidAuth)
	ro.e.POST("/api/v1/commentaries", c.CreateCommentary, middlewares.ValidAuth)
}
