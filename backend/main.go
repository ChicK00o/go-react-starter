package main

import (
	"backend/blackboard"
	"backend/config"
	"backend/db"
	"backend/log"
	"os"
	"os/signal"
)

func main() {
	closeChan := make(chan bool)
	log2 := log.NewLogger(false)
	defer log2.Close()

	app := NewApplication(log2)

	board := blackboard.NewBlackboard(app.log)
	utilities := NewUtilities(app.log)
	db2 := db.NewDatabase()
	defer db2.Close()
	con := config.NewConfig(log2, db2, "go-react-starter")

	setUpCloseListener(closeChan, log2, db2)
	routerConstruct := NewRouterConstruct(app.log, utilities, board, con, closeChan)
	routerConstruct.startRouter(con.Data.Port)
}

func setUpCloseListener(closeChan chan bool, logger log.Logger, database *db.Database) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	cleanup := func() {
		database.Close()
		logger.Close()
		os.Exit(3)
	}

	go func() {
		select {
		case <-c:
			cleanup()
		case <-closeChan:
			signal.Stop(c)
			cleanup()
		}
	}()

}
