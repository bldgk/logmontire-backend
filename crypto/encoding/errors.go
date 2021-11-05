package encoding

import (
	"fmt"
	"logmontire/common/errors"
	"net/http"
)

func JSONParseError(argument string) *errors.Error {
	return &errors.Error{
		Code: http.StatusBadRequest,
		What: fmt.Sprintf("Invalid json argument", argument),
	}
}
