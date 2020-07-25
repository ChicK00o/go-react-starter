package main

import (
	"4d63.com/tz"
	"backend/application"
	"backend/router"
	"github.com/ChicK00o/container"
	"time"
)

func main() {
	setIndianTimeZone()

	var app *application.Application
	container.Make(&app)
	defer app.Close()

	customCodeStart()

	var r *router.RouterConstruct
	container.Make(&r)
	r.StartRouter()
}

func setIndianTimeZone() {
	var timezone = "Asia/Kolkata"
	loc, err := tz.LoadLocation(timezone)
	if err != nil {
		panic(err.Error())
	}
	time.Local = loc
}
