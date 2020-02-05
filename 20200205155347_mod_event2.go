package main

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20200205155347, Down20200205155347)
}

func Up20200205155347(tx *sql.Tx) error {
	var err error
	_, err = tx.Exec("alter table event add column column2 char(255) not null")
	if err != nil {
		return err
	}

	return nil
}

func Down20200205155347(tx *sql.Tx) error {
	var err error
	_, err = tx.Exec("alter table event drop column column2")
	if err != nil {
		return err
	}
	return nil
}
