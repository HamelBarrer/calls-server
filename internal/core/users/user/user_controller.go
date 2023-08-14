package user

import (
	"fmt"
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

func (co *Controller) GetUser(c echo.Context) error {
	ui, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	u, err := co.r.GetById(ui)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, u)
}

func (co *Controller) GetUsers(c echo.Context) error {
	us, err := co.r.GetAllUser()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	if len(*us) == 0 {
		return echo.NewHTTPError(http.StatusNoContent, utils.ErrorHttp{Message: ""})
	}

	return c.JSON(http.StatusOK, us)
}

func (co *Controller) CreateUser(c echo.Context) error {
	u := new(UserCreate)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	m, e := utils.ValidEmpty[UserCreate](*u)
	if e {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: m})
	}

	if u.Password != u.PasswordConfirm {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: "The password not equals"})
	}

	nu, err := co.r.Create(*u)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, nu)
}

func (co *Controller) UpdatedUser(c echo.Context) error {
	ui, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	avatar, err := c.FormFile("avatar")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	location := fmt.Sprintf("files/images/%d/avatar", ui)
	file := fmt.Sprintf("files/images/%d/avatar/%s", ui, avatar.Filename)

	if err := utils.UploadFile(location, file, avatar); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	uu, err := co.r.Update(ui, file)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, uu)
}
