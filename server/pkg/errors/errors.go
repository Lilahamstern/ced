package errors

import (
	"net/http"
)

type (
	Op string

	Error struct {
		Kind int
		Op   Op
		Err  error
	}
)

const (
	KindNotFound       int = http.StatusNotFound
	KindInternalServer int = http.StatusInternalServerError
	KindAlreadyExists  int = http.StatusConflict
	KindValidation     int = http.StatusBadRequest
)

func (e *Error) Error() string {
	return ""
}

func E(args ...interface{}) *Error {
	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case int:
			e.Kind = arg
		case Op:
			e.Op = arg
		case error:
			e.Err = arg
		}
	}
	return e
}

func Kind(err error) int {
	e, ok := err.(*Error)
	if !ok {
		return KindInternalServer
	}

	if e.Kind != 0 {
		return e.Kind
	}

	return Kind(e.Err)
}

func Operations(e error) []Op {
	err, ok := e.(*Error)
	if !ok {
		return nil
	}

	res := []Op{err.Op}

	subErr, ok := e.(*Error)

	res = append(res, Operations(subErr)...)
	return res
}
