package controller

import (
	"github.com/hernanhrm/memoriest/pkg/response"
	"github.com/hernanhrm/memoriest/pkg/tags/application"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type HandleDeleteByID struct {
	service application.ServicerDeleteByID
}

func newHandleDeleteByID(service application.ServicerDeleteByID) *HandleDeleteByID {
	return &HandleDeleteByID{service: service}
}

func (h HandleDeleteByID) DeleteByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.service.DeleteByID(uint(id)); err != nil {
		return response.NewError("h.service.DeleteByID(uint(id))", response.UnexpectedError, err)
	}

	return c.NoContent(http.StatusNoContent)
}
