package mongo

import (
	"fmt"
	"logmontire/common/errors"
	"net/http"
)

func ObjectAlreadyExistsError(objectName string) *errors.Error {
	return &errors.Error{
		Code: http.StatusBadRequest,
		What: fmt.Sprintf("Object %s already exist", objectName),
	}
}

func ObjectNotExistsError(objectName string) *errors.Error {
	return &errors.Error{
		Code: http.StatusBadRequest,
		What: fmt.Sprintf("Object %s not exist", objectName),
	}
}

func InvalidObjectIdError(objectId string) *errors.Error {
	return &errors.Error{
		Code: http.StatusBadRequest,
		What: fmt.Sprintf("Invalid object id %s", objectId),
	}
}

func EmptyObjectId() *errors.Error {
	return &errors.Error{
		Code: http.StatusBadRequest,
		What: fmt.Sprint("Empty object id"),
	}
}

func MongoDBError(b error) *errors.Error {
	return &errors.Error{
		Code: http.StatusInternalServerError,
		What: fmt.Sprintf(b.Error()),
	}
}
