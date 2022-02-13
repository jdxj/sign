package sign_error

import (
	"fmt"
)

func New(code int, format string, a ...interface{}) error {
	se := &SignError{
		code: code,
		msg:  fmt.Sprintf(format, a...),
	}
	return se
}

func Wrap(code int, err error, format string, a ...interface{}) error {
	se := &SignError{
		code: code,
		msg:  fmt.Sprintf(format, a...),
		wrap: err,
	}
	return se
}

type SignError struct {
	code int
	msg  string
	wrap error
}

func (se *SignError) Error() string {
	var msg []string
	if se.msg != "" {
		msg = append(msg, se.msg)
	}
	if se.wrap != nil && se.wrap.Error() != "" {
		msg = append(msg, se.wrap.Error())
	}
	switch len(msg) {
	case 1:
		return msg[0]
	case 2:
		return fmt.Sprintf("%s, %s", msg[0], msg[1])
	}
	return ""
}
