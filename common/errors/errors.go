package errors

import (
	"fmt"
	"net/http"
)

func InvalidArgumentError(argument string) *Error {
	return &Error{
		Code: http.StatusBadRequest,
		What: fmt.Sprintf("Invalid argument %s", argument),
	}
}
