package validation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/thedevsaddam/govalidator"
)

var db *sqlx.DB

func init() {
	govalidator.AddCustomRule("unique", unique)
}

func unique(field string, rule string, message string, value interface{}) error {
	f := strings.Split(rule, ":")
	if len(f) < 2 {
		return errors.New("table is not specified on validation parameters")
	}

	exists, err := UniqueRecord(f[1], field, value)
	if err != nil {
		log.Warn(err)
		return errors.New("failed to validate unique record")
	}

	if exists {
		if message != "" {
			return errors.New(message)
		}
		return fmt.Errorf("the %s of %v already exists", field, value)
	}
	return nil
}

// RegisterValidation : Database instace
func RegisterValidation(DB *sqlx.DB) {
	db = DB
}

// UniqueRecord : Check if `data` by `field` and `table` does not exists.
func UniqueRecord(table string, field string, data interface{}) (bool, error) {
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s=$1)", table, field)
	stmt, err := db.Preparex(query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var exists bool
	err = stmt.Get(&exists, data)
	if err != nil {
		return false, err
	}

	return exists, nil
}
