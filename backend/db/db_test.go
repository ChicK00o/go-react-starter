package db

import (
	"context"
	"github.com/ChicK00o/container"
	"testing"
)

func TestDb (t *testing.T) {
	var db *Database
	container.Make(&db)

	defer db.Close()

	var greeting string
	err := db.Conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		t.Error("QueryRow failed: ", err)
	}
	t.Log(greeting)
}
