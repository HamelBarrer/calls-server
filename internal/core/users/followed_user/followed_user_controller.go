package followeduser

import (
	"net/http"
	"strconv"

	"github.com/HamelBarrer/calls-server/internal/utils"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	r Repository
}

func NewController(r Repository) *Controller {
	return &Controller{r}
}

func (co *Controller) GetFollowedUser(c echo.Context) error {
	fi, err := strconv.Atoi(c.Param("followed_user_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	f, err := co.r.GetById(fi)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, f)
}

func (co *Controller) GetFollowedUsers(c echo.Context) error {
	fs, err := co.r.GetAllFollowedUser()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	if len(*fs) == 0 {
		return echo.NewHTTPError(http.StatusNoContent, utils.ErrorHttp{Message: ""})
	}

	return c.JSON(http.StatusOK, fs)
}

func (co *Controller) CreateFollowedUser(c echo.Context) error {
	f := new(FollowedUser)
	if err := c.Bind(f); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	m, e := utils.ValidEmpty[FollowedUser](*f)
	if e {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: m})
	}

	fn, err := co.r.Create(*f)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, fn)
}
