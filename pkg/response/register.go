package response

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	codeErr := UnexpectedError

	if he, ok := err.(*echo.HTTPError); ok {
		msgErr, ok := he.Message.(string)
		if !ok {
			msgErr = "¡Upps! algo inesperado ocurrió"
		}

		registryError(c, he.Code, err)
		err = c.JSON(he.Code, Message{
			Error: &Errors{Message: msgErr}},
		)

		return
	}

	if e, ok := err.(*Error); ok {
		codeErr = e.Code
	}

	status, resp := getResponseError(codeErr)
	registryError(c, status, err)

	err = c.JSON(status, resp)

}

func registryError(c echo.Context, status int, err error) {
	fields := logrus.Fields{
		"Status":     status,
		"Enpoint":    c.Path(),
		"QueryParam": c.QueryParams(),
	}

	if e, ok := err.(*Error); ok {
		fields["Where"] = e.where
		fields["Who"] = e.who
	}

	if status >= 500 {
		logrus.WithFields(fields).Error(err)
		return
	}

	logrus.WithFields(fields).Warning(err)
}
