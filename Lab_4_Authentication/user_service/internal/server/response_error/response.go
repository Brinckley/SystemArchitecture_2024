package response_error

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
)

type Error struct {
	Err        error
	statusCode int
	message    string
	caller     string
}

func New(err error, statusCode int, message string) *Error {
	return &Error{
		Err:        err,
		statusCode: statusCode,
		message:    message,
		caller:     getCaller(),
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
		caller:     getCaller(),
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
		caller:     getCaller(),
	}
}

func (e *Error) Error() string {
	if e.message == "" {
		return e.Err.Error()
	}

	return fmt.Sprintf("%v\n[%v] > %v", e.Err, e.caller, e.message)
}

func (e *Error) Unwrap() error {
	return e.Err
}

func (e *Error) StatusCode() int {
	return e.statusCode
}

func getCaller() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}

	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short

	return fmt.Sprintf("%s:%d", file, line)
}
