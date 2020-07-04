package config

import (
	"backend/db"
	"backend/log"
	"context"
	jsoniter "github.com/json-iterator/go"
)

type Config struct {
	log  log.Logger
	db   *db.Database
	key  string
	Data *Data
}

type Data struct {
	Port        int    `json:"port"`
}

var configInstance *Config = nil

func NewConfig(l log.Logger, d *db.Database, projectKey string) *Config {
	if configInstance == nil {
		configInstance = &Config{
			log: l,
			db:  d,
			key: projectKey,
			Data: &Data{
				Port:        4999,
			},
		}
		configInstance.init()
	}
	return configInstance
}

func (c *Config) init() {
	//c.createTableIfNeeded()
	c.getConfigData()
}

func (c *Config) createTableIfNeeded() error {
	sqlStatement := `
		CREATE TABLE IF NOT EXISTS config
		(
			project_key TEXT PRIMARY KEY ,
			data json
		);`
	_, err := c.db.Conn.Exec(context.Background(), sqlStatement)
	if err != nil {
		c.log.Error(err)
		return err
	}
	return nil
}

func (c *Config) getConfigData() {
	val := &Data{}
	err := c.db.Conn.QueryRow(context.Background(), "SELECT data from config where project_key like $1", c.key).Scan(&val)
	if err != nil {
		c.log.Error(err)
		err = c.insertConfigData()
		if err != nil {
			c.log.Error(err)
		}
		return
	}
	c.Data = val
}

func (c *Config) insertConfigData() error {
	sqlStatement := `insert into config(project_key, data) values ($1, $2) on conflict (project_key) do update set data=$3`
	_, err := c.db.Conn.Exec(context.Background(), sqlStatement, c.key, c.Data, c.Data)
	if err != nil {
		c.log.Error(err)
		return err
	}
	return nil
}

func (c *Config) SaveConfigData() error {
	sqlStatement := `update config set data=$1 where project_key like $2`
	_, err := c.db.Conn.Exec(context.Background(), sqlStatement, c.Data, c.key)
	if err != nil {
		c.log.Error(err)
		return err
	}
	return nil
}

func (c *Config) UpdateConfig(msg interface{}) {
	data, err := jsoniter.ConfigFastest.Marshal(msg)
	if err != nil {
		c.log.Error(err)
		return
	}
	con := &Data{}
	err = jsoniter.ConfigFastest.Unmarshal(data, &con)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.Data = con
	err = c.SaveConfigData()
	if err != nil {
		c.log.Error(err)
		return
	}
}
