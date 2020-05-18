package response

import (
	"fmt"
	"net/http"
	"runtime"
)

// Error records an error, the http code and the
// message that caused it.
type Error struct {
	Code  CodeErr
	err   error
	where string
	who   string
}

type Message struct {
	Data  interface{} `json:"data,omitempty"`
	Error *Errors     `json:"error,omitempty"`
}

type Errors struct {
	Code    CodeErr `json:"code,omitempty"`
	Message string  `json:"message,omitempty"`
}

// NewError returns a rest.Error
func NewError(who string, co CodeErr, err error) *Error {
	fun, _, line, _ := runtime.Caller(1)

	return &Error{
		Code:  co,
		err:   err,
		where: fmt.Sprintf("%s:%d", runtime.FuncForPC(fun).Name(), line),
		who:   who,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf(e.err.Error())
}

// GetResponseError returns the status code and a ResponseError
func getResponseError(err CodeErr) (status int, resp Message) {
	status = http.StatusInternalServerError
	resErr := &Errors{err, "¡Upps! algo inesperado ocurrió"}

	if e, ok := errors[err]; ok {
		status = e.Status
		resErr.Message = e.Message
	}

	resp.Error = resErr
	return
}
