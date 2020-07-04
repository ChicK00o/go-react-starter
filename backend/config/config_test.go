package config

import (
	db2 "backend/db"
	log2 "backend/log"
	"testing"
)

func TestConfig(t *testing.T) {
	log := log2.Instance()
	db := db2.NewDatabase()
	config := NewConfig(log, db, "test1")

	t.Log(config.Data)

	t.Log(config.Data.AccessToken)

	config.Data.AccessToken = "testing5"

	err := config.SaveConfigData()
	if err != nil {
		t.Error(err)
	}
}
