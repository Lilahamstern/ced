package errors

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type (
	Op string

	Error struct {
		op     Op
		kind   string
		status int
		err    error
		level  logrus.Level
		data   interface{}
	}
)

const (
	KindSuccess string = "success"
	KindFail    string = "fail"
	KindError   string = "error"
)

func (e *Error) Error() string {
	if e.err != nil {
		return e.err.Error()
	}
	return ""
}

func (e *Error) Ops() []Op {
	return Ops(e)
}

func (e *Error) Kind() string {
	return Kind(e)
}

func (e *Error) Status() int {
	return Status(e)
}

func (e *Error) Data() interface{} {
	return Data(e)
}

func E(args ...interface{}) *Error {
	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case string:
			e.kind = arg
		case Op:
			e.op = arg
		case int:
			e.status = arg
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

	if e.data == nil {
		return "Message not implemented"
	}

	return e.data
}

func Kind(err error) string {
	e, ok := err.(*Error)
	if !ok {
		return KindError
	}

	if e.kind != "" {
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

func Status(err error) int {
	e, ok := err.(*Error)
	if !ok {
		return http.StatusInternalServerError
	}

	return e.status
}

func Level(err error) logrus.Level {
	e, ok := err.(*Error)
	if !ok {
		return 0
	}

	return e.level
}
