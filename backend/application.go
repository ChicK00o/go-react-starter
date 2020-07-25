package main

import (
	"backend/db"
	"backend/log"
	"github.com/ChicK00o/container"
	"os"
	"os/signal"
)

type Application struct {
	log       log.Logger
	db        *db.Database
	closeChan chan bool
}

func init() {
	container.Singleton(func(logger log.Logger, database *db.Database) *Application {
		app := &Application{
			log:       logger,
			db:        database,
			closeChan: make(chan bool),
		}
		app.setUpCloseListener()
		return app
	})
}

func (app *Application) setUpCloseListener() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	cleanup := func() {
		app.Close()
		os.Exit(3)
	}

	go func() {
		select {
		case <-c:
			cleanup()
		case <-app.closeChan:
			signal.Stop(c)
			cleanup()
		}
	}()
}

func (app *Application) Close() {
	app.db.Close()
	app.log.Close()
}
