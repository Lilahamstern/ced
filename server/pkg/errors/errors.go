package errors

import (
	"net/http"
)

type (
	// Op : Operation string for custom stack traces
	Op     string
	Status string
	// Error : Error struct for custom errors better logging
	Error struct {
		op     Op
		status Status
		code   int
		err    error
		data   interface{}
	}
)

const (
	// KindSuccess : Code success
	KindSuccess Status = "success"
	// KindFail : Code fail
	KindFail Status = "fail"
	// KindError : Code error
	KindError Status = "error"
)

func (e *Error) Error() string {
	if e.err != nil {
		return e.err.Error()
	}

	return ""
}

// Ops : Get operations in error
func (e *Error) Ops() []Op {
	return Ops(e)
}

// Status : Get error kind
func (e *Error) Status() Status {
	return status(e)
}

// Code : Get error code
func (e *Error) Code() int {
	return Code(e)
}

// Data : Get error data
func (e *Error) Data() interface{} {
	return Data(e)
}

// E : Generate new error
func E(args ...interface{}) *Error {
	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case Status:
			e.status = arg
		case Op:
			e.op = arg
		case int:
			e.code = arg
		case error:
			e.err = arg
		case interface{}:
			e.data = arg
		}
	}
	return e
}

// Data : Get error data
func Data(err error) interface{} {
	e, ok := err.(*Error)
	if !ok {
		return "Internal server error"
	}

	if e.data != nil {
		return e.data
	}

	return Data(e.err)
}

func status(err error) Status {
	e, ok := err.(*Error)
	if !ok {
		return KindError
	}

	if e.status != "" {
		return e.status
	}

	return status(e.err)
}

// Ops : Get error operations
func Ops(e *Error) []Op {
	res := []Op{e.op}

	subErr, ok := e.err.(*Error)
	if !ok {
		return res
	}

	res = append(res, Ops(subErr)...)

	return res
}

// Code : Get error code
func Code(err error) int {
	e, ok := err.(*Error)
	if !ok {
		return http.StatusInternalServerError
	}

	return e.code
}

func (s Status) String() string {
	return string(s)
}
