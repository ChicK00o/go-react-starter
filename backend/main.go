package main

import (
	"4d63.com/tz"
	"github.com/ChicK00o/container"
	"time"
)

func main() {
	setIndianTimeZone()

	var app *Application
	container.Make(&app)
	defer app.Close()

	customCodeStart()

	var router *RouterConstruct
	container.Make(&router)
	router.StartRouter()
}

func setIndianTimeZone() {
	var timezone = "Asia/Kolkata"
	loc, err := tz.LoadLocation(timezone)
	if err != nil {
		panic(err.Error())
	}
	time.Local = loc
}
