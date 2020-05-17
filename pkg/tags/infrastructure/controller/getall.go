package controller

import (
	"encoding/json"
	dsl "github.com/hernanhrm/dslparser"
	"github.com/hernanhrm/memoriest/pkg/response"
	"github.com/hernanhrm/memoriest/pkg/tags/application"
	"github.com/labstack/echo"
	"net/http"
)

type HandleGetAll struct {
	servicer application.ServicerGetAll
}

func NewHandleGetAll(servicer application.ServicerGetAll) *HandleGetAll {
	return &HandleGetAll{servicer: servicer}
}

func (h HandleGetAll) GetAll(c echo.Context) error {
	dslModel := dsl.Model{}
	query := c.QueryParam("query")

	if err := json.Unmarshal([]byte(query), &dslModel); err != nil {
		return response.NewError("json.Unmarshal([]byte(query), &dslModel)", response.BindFailed, err)
	}

	data, err := h.servicer.GetAll(dslModel)
	if err != nil {
		return response.NewError("h.servicer.GetAll(dslModel)", response.UnexpectedError, err)
	}

	return c.JSON(http.StatusOK, data)
}
