package config

import (
	"github.com/ChicK00o/container"
	"testing"
)

func TestConfig(t *testing.T) {
	var config *Config
	container.Make(&config)
	defer config.db.Close()

	t.Log(config.Data)

	err := config.SaveConfigData()
	if err != nil {
		t.Error(err)
	}
}
