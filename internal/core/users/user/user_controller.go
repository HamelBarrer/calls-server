package user

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
	userId, err := strconv.Atoi(c.Get("userId").(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	us, err := co.r.GetAllUser(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	if len(*us) == 0 {
		return echo.NewHTTPError(http.StatusNoContent, utils.ErrorHttp{Message: ""})
	}

	return c.JSON(http.StatusOK, us)
}

func (co *Controller) CreateUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	passwordConfirm := c.FormValue("passwordConfirm")

	u := UserCreate{
		User:            User{Username: username},
		Password:        password,
		PasswordConfirm: passwordConfirm,
	}

	if u.Password != u.PasswordConfirm {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: "The password not equals"})
	}

	nu, err := co.r.Create(u)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	ubicationFile, err := utils.UploadFile("avatar", nu.UserId, c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	uu, err := co.r.Update(nu.UserId, ubicationFile)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, uu)
}

func (co *Controller) UpdatedUser(c echo.Context) error {
	ui, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	ubicationFile, err := utils.UploadFile("avatar", ui, c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	uu, err := co.r.Update(ui, ubicationFile)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, uu)
}
