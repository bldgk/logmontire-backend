package http

import (
	"logmontire/common/errors"
)

type Response struct {
	Body  []byte
	Code  int
	Error *errors.Error
}

func (resp *Response) SetBody(body []byte) {
	resp.Body = body
}

func (resp *Response) SetBodyString(body string) {
	resp.Body = []byte(body)
}

func (resp *Response) SetError(err *errors.Error) {
	resp.Error = err
	resp.SetBodyString(err.Error())
	resp.Code = err.Code
}
