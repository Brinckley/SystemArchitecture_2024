package response_error

import (
	"errors"
	"fmt"
	"net/http"
)

type Error struct {
	Err        error
	statusCode int
	message    string
}

func New(err error, statusCode int, message string) *Error {
	return &Error{
		Err:        err,
		statusCode: statusCode,
		message:    message,
	}
}

func Wrap(err error, msg string) error {
	var e *Error
	statusCode := http.StatusInternalServerError

	if errors.As(err, &e) {
		statusCode = e.statusCode
	}

	return &Error{
		Err:        err,
		statusCode: statusCode,
		message:    msg,
	}
}

func From(statusCode int) error {
	text := http.StatusText(statusCode)

	if text == "" {
		text = http.StatusText(http.StatusInternalServerError)
		statusCode = http.StatusInternalServerError
	}

	return &Error{
		Err:        errors.New(text),
		statusCode: statusCode,
		message:    "",
	}
}

func (e *Error) Message() string {
	if e.message != "" {
		return e.message
	}
	return fmt.Sprintf("no msg for this error %v", e.Err)
}

func (e *Error) Error() string {
	if e.message == "" {
		return e.Err.Error()
	}

	return fmt.Sprintf("%v > %v", e.Err, e.message)
}

func (e *Error) Unwrap() error {
	return e.Err
}

func (e *Error) StatusCode() int {
	return e.statusCode
}
