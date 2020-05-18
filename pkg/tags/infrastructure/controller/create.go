package controller

import (
	"github.com/hernanhrm/memoriest/pkg/response"
	"github.com/hernanhrm/memoriest/pkg/tags/application"
	"github.com/hernanhrm/memoriest/pkg/tags/domain"
	"github.com/labstack/echo"
	"net/http"
)

type HandleCreate struct {
	service application.ServicerCreate
}

func newHandleCreate(service application.ServicerCreate) *HandleCreate {
	return &HandleCreate{service: service}
}

func (h HandleCreate) Create(c echo.Context) error {
	model := domain.Command{}

	if err := c.Bind(&model); err != nil {
		return response.NewError("c.Bind(&model)", response.BindFailed, err)
	}

	if err := h.service.Create(&model); err != nil {
		return response.NewError("h.service.Create(&model)", response.UnexpectedError, err)
	}

	return c.JSON(http.StatusCreated, model)
}
