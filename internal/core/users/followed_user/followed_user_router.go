package followeduser

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

	ro.e.GET("/api/v1/followed_user/:followed_user_id", c.GetFollowedUser)
	ro.e.GET("/api/v1/followed_user", c.GetFollowedUsers)
	ro.e.POST("/api/v1/followed_user", c.CreateFollowedUser)
}
