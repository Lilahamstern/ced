package logger

import (
	"github.com/lilahamstern/ced/server/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func SystemErr(err error) {
	sysErr, ok := err.(*errors.Error)
	if !ok {
		Error(err)
		return
	}

	entry := log.WithFields(log.Fields{
		"operations": errors.Ops(sysErr),
		"kind":       errors.Kind(sysErr),
		"data":       errors.Data(sysErr),
	})

	switch errors.Level(err) {
	case log.WarnLevel:
		entry.Warnf("%s: %v", sysErr.Ops(), err)
	case log.InfoLevel:
		entry.Infof("%s: %v", sysErr.Ops(), err)
	case log.DebugLevel:
		entry.Debugf("%s: %v", sysErr.Ops(), err)
	default:
		entry.Errorf("%s: %v", sysErr.Ops(), err)
	}
}

func Error(err error) {
	log.Error(err)
}
