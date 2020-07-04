package db

import (
	"context"
	"testing"
)

func TestDb (t *testing.T) {
	db := NewDatabase()
	defer db.Close()

	var greeting string
	err := db.Conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		t.Error("QueryRow failed: ", err)
	}
	t.Log(greeting)

}
