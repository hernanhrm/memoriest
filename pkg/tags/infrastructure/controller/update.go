package controller

import (
	"github.com/hernanhrm/memoriest/pkg/response"
	"github.com/hernanhrm/memoriest/pkg/tags/application"
	"github.com/hernanhrm/memoriest/pkg/tags/domain"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type HandleUpdateByID struct {
	service application.ServicerUpdateByID
}

func newHandleUpdateByID(service application.ServicerUpdateByID) *HandleUpdateByID {
	return &HandleUpdateByID{service: service}
}

func (h HandleUpdateByID) Update(c echo.Context) error {
	model := domain.Command{}

	if err := c.Bind(&model); err != nil {
		return response.NewError("c.Bind(&model)", response.BindFailed, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	model.ID = uint(id)

	if err := h.service.UpdateByID(model); err != nil {
		return response.NewError("h.service.UpdateByID(model)", response.UnexpectedError, err)
	}

	return c.NoContent(http.StatusOK)
}
