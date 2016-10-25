package models

import "fmt"

type ModelError struct {
	HTTPCode     int
	ErrorMessage string
}

func (me *ModelError) Error() string {
	return me.ErrorMessage
}

func NewModelError(code int, messagefmt string, args ...interface{}) *ModelError {
	return &ModelError{
		HTTPCode:     code,
		ErrorMessage: fmt.Sprintf(messagefmt, args...),
	}
}
