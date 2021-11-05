package crypto

import (
	"fmt"
	"logmontire/common/errors"
)

func InvalidSignatureError() *errors.Error {
	return &errors.Error{
		Code: 601,
		What: fmt.Sprintf("Invalid signature"),
	}
}

func InvalidHashError() *errors.Error {
	return &errors.Error{
		Code: 602,
		What: fmt.Sprintf("Invalid hash"),
	}
}
