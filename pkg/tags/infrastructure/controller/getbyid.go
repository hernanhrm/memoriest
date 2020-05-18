package controller

import (
	"encoding/json"
	dsl "github.com/hernanhrm/dslparser"
	"github.com/hernanhrm/memoriest/pkg/response"
	"github.com/hernanhrm/memoriest/pkg/tags/application"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type HandleGetByID struct {
	service application.ServicerGetByID
}

func newHandleGetByID(service application.ServicerGetByID) *HandleGetByID {
	return &HandleGetByID{service: service}
}

func (h HandleGetByID) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	dslModel := dsl.Model{}
	query := c.QueryParam("query")

	if err := json.Unmarshal([]byte(query), &dslModel); err != nil {
		return response.NewError("json.Unmarshal([]byte(query), &dslModel)", response.BindFailed, err)
	}

	data, err := h.service.GetByID(uint(id), dslModel)
	if err != nil {
		return response.NewError("h.service.GetByID(uint(id)", response.UnexpectedError, err)
	}

	return c.JSON(http.StatusOK, data)
}
