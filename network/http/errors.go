package http

import (
	"fmt"
	"logmontire/common/errors"
	"net/http"
)

func RouteNotFindError(path string) *errors.Error {
	return &errors.Error{
		Code: http.StatusNotFound,
		What: fmt.Sprintf("Route %s not found", path),
	}
}

func MethodNotAllowedError(method string) *errors.Error {
	return &errors.Error{
		Code: http.StatusMethodNotAllowed,
		What: fmt.Sprintf("Method %s not allowed", method),
	}
}

func BadRequestError() *errors.Error {
	return &errors.Error{
		Code: http.StatusBadRequest,
		What: fmt.Sprintf("Bad request"),
	}
}

func ConnectionRefusedError(url string) *errors.Error {
	return &errors.Error{
		Code: http.StatusInternalServerError,
		What: fmt.Sprintf("Connection refused to " + url),
	}
}
