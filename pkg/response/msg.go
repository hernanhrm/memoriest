package response

import "net/http"

// CodeErr type errors of the API
type CodeErr int

// errorTypes
const (
	UnexpectedError CodeErr = iota + 1
	UniqueError
	NotNullError
	InvalidParameter
	BindFailed
	ServiceUnauthorized
	BulkError
	BuildQueryError
	NotFound
)

type dataError struct {
	Status  int
	Message string
}

var errors = map[CodeErr]dataError{
	UniqueError: {
		Status:  http.StatusBadRequest,
		Message: "¡Upps! nos enviaste un registro duplicado, revisa la documentación del paquete",
	},
	NotNullError: {
		Status:  http.StatusBadRequest,
		Message: "¡Upps! nos enviaste datos nulos, revisa la documentación del paquete",
	},
	InvalidParameter: {
		Status:  http.StatusBadRequest,
		Message: "¡Upps! el valor del parámetro que has enviado no es valido. Por favor revisa el formato y tipo",
	},
	BindFailed: {
		Status:  http.StatusBadRequest,
		Message: "¡Upps! el payload que enviaste no es valido, verifica la documentación del paquete",
	},
	ServiceUnauthorized: {
		Status:  http.StatusUnauthorized,
		Message: "¡Upps! no estas autorizado para consumir este servicio",
	},
	BulkError: {
		Status:  http.StatusBadRequest,
		Message: "¡Upps! no se pudo bolcar la data",
	},
	BuildQueryError: {
		Status:  http.StatusInternalServerError,
		Message: "¡Upps! no se pudo crear el request body",
	},
	NotFound: {
		Status:  http.StatusNotFound,
		Message: "¡Upps! hubo un problema al buscar el documento",
	},
}
