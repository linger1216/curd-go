package core

import "net/http"

var (
	StatusInternalServerError = NewError(http.StatusInternalServerError, "Status Internal Server Error")
)

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) StatusCode() int {
	return e.Code
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
