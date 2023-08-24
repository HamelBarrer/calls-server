package middlewares

import (
	"net/http"
	"strings"

	"github.com/HamelBarrer/calls-server/internal/utils"
	"github.com/labstack/echo/v4"
)

func ValidAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		a := c.Request().Header.Get("Authorization")
		if a == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrorHttp{Message: "Information not found"})
		}

		au := strings.Split(a, " ")
		if len(au) != 2 {
			return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrorHttp{Message: "Information not found"})
		}

		t := au[1]
		j, v, err := utils.ValidateToken(t)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrorHttp{Message: err.Error()})
		}

		if !v {
			return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrorHttp{Message: "Information not found"})
		}

		c.Set("userId", j.UserId)

		return next(c)
	}
}
