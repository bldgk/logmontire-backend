package errors

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"time"
)

type Error struct {
	What string `json:"error"`
	Code int    `json:"code"`
}

func (e *Error) Error() string {
	r := color.New(color.FgRed, color.Bold)
	r.Println(fmt.Sprintf("Error at %v: %s", time.Now(), e.What))
	js, _ := json.Marshal(e)
	return string(js)
}

func IsError(errs ...*Error) *Error {
	for _, err := range errs {
		if err != nil {
			//err.Error()
			return err
		}
	}
	return nil
}
