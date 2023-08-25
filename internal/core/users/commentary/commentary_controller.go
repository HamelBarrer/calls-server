package commentary

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

func (co *Controller) GetCommentary(c echo.Context) error {
	ci, err := strconv.Atoi(c.Param("commentary_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	com, err := co.r.GetById(ci)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, com)
}

func (co *Controller) GetCommentaries(c echo.Context) error {
	userId, err := strconv.Atoi(c.Get("userId").(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	cs, err := co.r.GetAll(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	if len(*cs) == 0 {
		return echo.NewHTTPError(http.StatusNoContent, utils.ErrorHttp{Message: ""})
	}

	return c.JSON(http.StatusOK, cs)
}

func (co *Controller) CreateCommentary(c echo.Context) error {
	userId, err := strconv.Atoi(c.Get("userId").(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	com := new(Commentary)
	if err := c.Bind(com); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	com.UserId = userId

	m, isEmpty := utils.ValidEmpty[Commentary](*com)
	if isEmpty {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: m})
	}

	cn, err := co.r.Create(com)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorHttp{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, cn)
}
