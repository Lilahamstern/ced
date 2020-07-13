package validation

import (
	"errors"
	"fmt"
	database "github.com/lilahamstern/ced/server/internal/db/postgres"
	"github.com/thedevsaddam/govalidator"
	"strings"
)

var session *database.Session

func init() {
	govalidator.AddCustomRule("unique", unique)
}

func unique(field string, rule string, message string, value interface{}) error {
	f := strings.Split(rule, ":")
	if len(f) < 2 {
		return errors.New("Table is not specified on validation parameters")
	}

	exists, err := session.UniqueRecord(f[1], field, value)
	if err != nil {
		return err
	}

	if exists {
		if message != "" {
			return errors.New(message)
		}
		return fmt.Errorf("The %s of %v already exists.", field, value)
	}
	return nil
}

func Register(s *database.Session) {
	session = s
}
