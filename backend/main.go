package main

import "backend/log"

func main() {
	log := log.NewLogger(false)
	app := NewApplication(log)
	defer func() {
		log.Close()
	}()

	blackboard := NewBlackboard(app.log)
	utilities := NewUtilities(app.log)

	routerConstruct := NewRouterConstruct(app.log, utilities, blackboard)
	routerConstruct.startRouter(5000)
}
