package util

import (
	"database/sql"

	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
	"github.com/hooligram/kifu"
)

func SpinUpTestDb() *sql.DB {
	db, err := mysqltestcontainer.Start("test", "./../migration")
	if err != nil {
		kifu.Fatal(err.Error())
		panic(err.Error())
	}
	return db
}
