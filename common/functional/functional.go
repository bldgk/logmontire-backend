package functional

import (
	"logmontire/common/errors"
)

type F func() *errors.Error

func Fs(fs ...F) *errors.Error {
	for _, f := range fs {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
