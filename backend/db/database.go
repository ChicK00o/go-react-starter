package db

import (
	"backend/log"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

type Database struct {
	Conn *pgxpool.Pool
}

var instance *Database

func NewDatabase() *Database {
	if instance == nil {
		instance = &Database{}
		logger := log.Instance()
		instance.connectToDb(logger)
	}
	return instance
}

func (d *Database) Close() {
	d.Conn.Close()
}

func (d *Database) connectToDb(logger log.Logger) {
	//   # Example DSN
	//   user=jack password=secret host=host1,host2,host3 port=5432,5433,5434 dbname=mydb sslmode=verify-ca
	dsn := "user=" + DatabaseUserName
	dsn += " " + "password=" + DatabasePassword
	dsn += " " + "host=" + DatabaseURL
	dsn += " " + "port=" + DatabasePort
	dsn += " " + "dbname=market-options"
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		logger.Panic(err.Error())
	}

	if log.IsZapLogger() {
		config.ConnConfig.Logger = zapadapter.NewLogger(log.GetZapLogger().Desugar())
		config.ConnConfig.LogLevel = pgx.LogLevelWarn
	}

	dbpool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		logger.Panic(err)
		os.Exit(1)
	}

	d.Conn = dbpool
	return
}

const (
	DatabaseURL string = "192.168.0.171"
	DatabasePort string = "5432"
	DatabaseUserName string = "postgres"
	DatabasePassword string = "itsmerohit"
)
