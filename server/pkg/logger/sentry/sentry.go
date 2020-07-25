package sentry

import (
	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
	"time"
)

func Register() (func(), error) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: viper.GetString("sentry.dsn"),
	})

	if err != nil {
		return nil, err
	}

	tidy := func() {
		sentry.Flush(2 * time.Second)
	}

	return tidy, nil
}
