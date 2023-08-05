package server

import (
	followeduser "github.com/HamelBarrer/calls-server/internal/core/users/followed_user"
	"github.com/HamelBarrer/calls-server/internal/core/users/user"
	"github.com/HamelBarrer/calls-server/internal/storage"
	"github.com/labstack/echo/v4"
)

func Handler(s storage.Storage) {
	e := echo.New()

	ur := user.Newrepository(s)
	uc := user.NewRouter(e)
	uc.SetupConfig(ur)

	fs := followeduser.NewService(s)
	fc := followeduser.NewRouter(e)
	fc.SetupConfig(fs)

	e.Logger.Fatal(e.Start(":3000"))
}
