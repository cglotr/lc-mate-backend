package util

import (
	"database/sql"

	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
	"github.com/hooligram/kifu"
)

func SpinUpTestDb(migrationDir string) *sql.DB {
	db, err := mysqltestcontainer.Start("test", migrationDir)
	if err != nil {
		kifu.Fatal(err.Error())
	}
	return db
}
