package main

import (
	"testing"

	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
)

func TestMigrate(t *testing.T) {
	db, err := mysqltestcontainer.Start("test", "./migration")
	if err != nil {
		t.Errorf("db error: %v\n", err.Error())
	}

	usernames := []string{"awice", "larryNY", "numb3r5"}
	for _, username := range usernames {
		db.Exec(`INSERT INTO user (username) VALUES (?);`, username)
	}

	rows, err := db.Query("SELECT id, username FROM user ORDER BY id ASC;")
	if err != nil {
		t.Errorf("db query error: %v\n", err.Error())
	}
	for i := 0; i < len(usernames); i++ {
		rows.Next()
		var id int
		var username string
		rows.Scan(&id, &username)
		if i+1 != id {
			t.Errorf("expected %v, got %v\n", i+1, id)
		}
		if usernames[i] != username {
			t.Errorf("expected %v, got %v\n", usernames[i], username)
		}
	}
}
