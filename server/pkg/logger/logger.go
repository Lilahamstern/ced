package logger

import (
	"github.com/lilahamstern/ced/server/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
)

func Register() {
	l := log.New()
	l.SetOutput(os.Stdout)
}

func SystemErr(err error) {
	e, ok := err.(*errors.Error)
	if !ok {
		Error(err)
		return
	}

	entry := log.WithFields(log.Fields{
		"operations": e.Ops(),
		"data":       e.Data(),
		"status":     e.Code(),
	})

	switch e.Kind() {
	case errors.KindSuccess:
		entry.Infof("%s: %v", e.Ops(), err)
	case errors.KindFail:
		entry.Warnf("%s: %v", e.Ops(), err)
	default:
		entry.Errorf("%s: %v", e.Ops(), err)
	}
}

func Error(err error) {
	log.Error(err)
}
