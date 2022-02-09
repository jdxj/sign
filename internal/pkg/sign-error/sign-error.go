package sign_error

import (
	"fmt"
)

func New(code int, format string, a ...interface{}) error {
	se := &SignError{
		Code:        code,
		CodeDesc:    ErrMap[code],
		Description: fmt.Sprintf(format, a...),
	}
	return se
}

type SignError struct {
	code int
	err  error
}

func (se *SignError) Error() string {
	return se.Description
}
