package errors

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func NewFromErr(err error) *Error {
	internalErr, ok := err.(*Error)
	if ok {
		return internalErr
	}

	return &Error{
		Type:    InternalError,
		Message: err.Error(),
		Stack:   debug.Stack(),
	}
}

func New(msg string) *Error {
	return newError(DataError, msg)
}
func Newf(format string, args ...interface{}) *Error {
	return newErrorf(DataError, format, args...)
}

func NewInternalError(msg string) *Error {
	return newError(InternalError, msg).AddCode(http.StatusInternalServerError)
}
func NewInternalErrorf(format string, args ...interface{}) *Error {
	return newErrorf(InternalError, format, args...)
}

func NewCriticalError(msg string) *Error {
	return newError(CriticalError, msg)
}
func NewCriticalErrorf(format string, args ...interface{}) *Error {
	return newErrorf(CriticalError, format, args...)
}

func newErrorf(t ErrorType, format string, args ...interface{}) *Error {
	return newError(t, fmt.Sprintf(format, args...))
}
func newError(t ErrorType, msg string) *Error {
	return &Error{
		Type:    t,
		Message: msg,
		Stack:   debug.Stack(),
	}
}

func (e *Error) AddCode(code int) *Error {
	e.Code = code

	return e
}
func (e *Error) AddError(err error) *Error {
	e.ExtendedError = err
	return e
}

func (e *Error) ConvertToClientError() *ClientError {
	return &ClientError{Message: e.Message}
}

func (e *Error) ConvertToTransportError() *ErrorTransport {
	return &ErrorTransport{
		Type:    e.Type,
		Message: e.Message,
		Code:    e.Code,
	}
}
func (e *ErrorTransport) ConvertToError() *Error {
	return newError(e.Type, e.Message).AddCode(e.Code)
}
