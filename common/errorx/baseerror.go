package errorx

import "fmt"

const defaultCode = 1001
const systemCode = 1002

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewDefaultError(msg ...string) error {
	return NewCodeError(defaultCode, fmt.Sprint(msg))
}
func NewSystemError(msg ...string) error {
	return NewCodeError(systemCode, fmt.Sprint(msg))
}
func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
