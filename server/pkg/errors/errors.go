package errors

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type (
	Op string

	Error struct {
		kind  int
		op    Op
		err   error
		level logrus.Level
		data  interface{}
	}
)

const (
	KindNotFound       int = http.StatusNotFound
	KindInternalServer int = http.StatusInternalServerError
	KindAlreadyExists  int = http.StatusConflict
	KindValidation     int = http.StatusBadRequest
)

func (e *Error) Error() string {
	if e.err != nil {
		return e.Error()
	}
	return ""
}

func (e *Error) Ops() []Op {
	return Ops(e)
}

func (e *Error) Data() interface{} {
	return Data(e)
}

func (e *Error) Kind() int {
	return Kind(e)
}

func E(args ...interface{}) *Error {
	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case int:
			e.kind = arg
		case Op:
			e.op = arg
		case error:
			e.err = arg
		case logrus.Level:
			e.level = arg
		case interface{}:
			e.data = arg
		}
	}
	return e
}

func Data(err error) interface{} {
	e, ok := err.(*Error)
	if !ok {
		return "Internal server error"
	}

	if e.data != nil {
		return e.data
	}

	return Data(e)
}

func Kind(err error) int {
	e, ok := err.(*Error)
	if !ok {
		return KindInternalServer
	}

	if e.kind != 0 {
		return e.kind
	}

	return Kind(e.err)
}

func Ops(e *Error) []Op {
	res := []Op{e.op}

	subErr, ok := e.err.(*Error)
	if !ok {
		return res
	}

	res = append(res, Ops(subErr)...)

	return res
}

func Level(err error) logrus.Level {
	e, ok := err.(*Error)
	if !ok {
		return 0
	}

	return e.level
}
