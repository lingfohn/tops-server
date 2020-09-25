package svcerr

import "fmt"

var (
	InternalServerError = &ErrNo{code: 1199999, format: "内部服务错误"}
	JSONBindErr         = NewSvcErrCode(1199800, "请求参数错误")
	CommonUnmarshal     = NewSvcErrCode(1199801, "反序列化yaml失败")
)

type ErrNo struct {
	code   int
	format string
}

func (e ErrNo) Code() int {
	return e.code
}

func NewSvcErrCode(code int, message string) *ErrNo {
	return &ErrNo{code: code, format: message}
}

type SvcError struct {
	*ErrNo
	message string
	Err     error
}

func (e *SvcError) Error() string {
	return e.message
}

func (e *SvcError) Unwrap() error {
	return e.Err
}

func Error(err error, no *ErrNo) error {
	return &SvcError{ErrNo: no, message: no.format, Err: err}
}

func Errorf(err error, no *ErrNo, a ...interface{}) error {
	return &SvcError{ErrNo: no, message: fmt.Sprintf(no.format, a...), Err: err}
}
