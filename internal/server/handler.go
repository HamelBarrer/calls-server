package server

import (
	"github.com/HamelBarrer/calls-server/internal/core/user"
	"github.com/HamelBarrer/calls-server/internal/storage"
	"github.com/labstack/echo/v4"
)

func Handler(s storage.Storage) {
	e := echo.New()

	ur := user.Newrepository(s)
	uc := user.NewRouter(e)
	uc.SetupConfig(ur)

	e.Logger.Fatal(e.Start(":3000"))
}
