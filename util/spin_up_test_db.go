package util

import (
	"database/sql"

	"github.com/arikama/go-arctic-tern/arctictern"
	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
	"github.com/hooligram/kifu"
)

func SpinUpTestDb(migrationDir string) *sql.DB {
	result, err := mysqltestcontainer.Start("test")
	if err != nil {
		kifu.Fatal(err.Error())
	}
	arctictern.Migrate(result.Db, migrationDir)
	return result.Db
}
