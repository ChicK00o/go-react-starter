package main

import (
	"4d63.com/tz"
	"backend/log"
	"time"
)

type Application struct {
	log   log.Logger
}

func NewApplication(log log.Logger) *Application {
	app := &Application{
		log:   log,
	}
	app.setIndianTimeZone()
	return app
}

func (app *Application) setIndianTimeZone() {
	var timezone = "Asia/Kolkata"
	loc, err := tz.LoadLocation(timezone)
	if err != nil {
		app.log.Panic(err.Error())
	}
	time.Local = loc
}
