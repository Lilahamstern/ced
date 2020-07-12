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
	}
)

const (
	KindNotFound       int = http.StatusNotFound
	KindInternalServer int = http.StatusInternalServerError
	KindAlreadyExists  int = http.StatusConflict
	KindValidation     int = http.StatusBadRequest
)

func (e *Error) Error() string {
	return e.err.Error()
}

func (e *Error) Ops() []Op {
	return Ops(e)
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
		}
	}
	return e
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
