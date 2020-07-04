package main

import (
	"backend/blackboard"
	"backend/config"
	"backend/db"
	"backend/log"
)

func main() {
	log := log.NewLogger(false)
	app := NewApplication(log)
	defer func() {
		log.Close()
	}()

	board := blackboard.NewBlackboard(app.log)
	utilities := NewUtilities(app.log)
	dbase := db.NewDatabase()
	con := config.NewConfig(log, dbase, "go-react-starter")

	routerConstruct := NewRouterConstruct(app.log, utilities, board, con)
	routerConstruct.startRouter(con.Data.Port)
}
