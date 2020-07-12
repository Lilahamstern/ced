package logger

import (
	"github.com/lilahamstern/ced/server/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
)

type Logger interface {
	SystemErr(err error)
	Error(err error)
}

type logger struct {
}

var log *logrus.Logger

func (l logger) SystemErr(err error) {
	sysErr, ok := err.(*errors.Error)
	if !ok {
		l.Error(err)
		return
	}

	entry := log.WithFields(logrus.Fields{
		"operations": errors.Ops(sysErr),
		"kind":       errors.Kind(sysErr),
	})

	switch errors.Level(err) {
	case logrus.WarnLevel:
		entry.Warnf("%s: %v", sysErr.Ops(), err)
	case logrus.InfoLevel:
		entry.Infof("%s: %v", sysErr.Ops(), err)
	case logrus.DebugLevel:
		entry.Debugf("%s: %v", sysErr.Ops(), err)
	default:
		entry.Errorf("%s: %v", sysErr.Ops(), err)
	}
}

func (l logger) Error(err error) {
	log.Error(err)
}

func New() Logger {
	log = logrus.New()
	log.Out = os.Stdout

	return &logger{}
}
