package customauth

import (
	"net/http"

	"github.com/HamelBarrer/calls-server/internal/utils"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	r Repository
}

func NewController(r Repository) *Controller {
	return &Controller{r}
}

func (co *Controller) Login(c echo.Context) error {
	au := new(Auth)
	if err := c.Bind(au); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	m, e := utils.ValidEmpty[Auth](*au)
	if e {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: m})
	}

	u, err := co.r.GetUserByUsername(au.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	eq, err := utils.VerifyHash(au.Password, u.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	if !eq {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: "Username or password incorrect"})
	}

	t, err := utils.CreationToken(u.UserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user_id":  u.UserId,
		"username": u.Username,
		"token":    t,
	})
}
