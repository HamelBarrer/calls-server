package server

import (
	customauth "github.com/HamelBarrer/calls-server/internal/core/auths/custom_auth"
	followeduser "github.com/HamelBarrer/calls-server/internal/core/users/followed_user"
	"github.com/HamelBarrer/calls-server/internal/core/users/user"
	"github.com/HamelBarrer/calls-server/internal/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Handler(s storage.Storage) {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	ur := user.Newrepository(s)
	uc := user.NewRouter(e)
	uc.SetupConfig(ur)

	cr := customauth.NewService(s)
	cc := customauth.NewRouter(e)
	cc.SetupConfig(cr)

	fs := followeduser.NewService(s)
	fc := followeduser.NewRouter(e)
	fc.SetupConfig(fs)

	e.Logger.Fatal(e.Start(":3000"))
}
