package main

import (
	"backend/blackboard"
	"backend/config"
	"backend/log"
	"backend/websocket"
	"fmt"
	"github.com/ChicK00o/container"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"time"
)

type RouterConstruct struct {
	log        log.Logger
	utilities  *Utilities
	router     *gin.Engine
	blackboard *blackboard.Blackboard
	pool       *websocket.Pool
	config     *config.Config
	closeChan  chan bool
}

func init() {
	gin.SetMode(gin.ReleaseMode)
	container.Singleton(func(
		l log.Logger,
		u *Utilities,
		b *blackboard.Blackboard,
		c *config.Config,
		p *websocket.Pool,
		app *Application) *RouterConstruct {
		return &RouterConstruct{
			log:        l,
			utilities:  u,
			router:     gin.Default(),
			blackboard: b,
			config:     c,
			pool:       p,
			closeChan:  app.closeChan}
	})
}

func (c *RouterConstruct) StartRouter() {
	// Set the router as the default one shipped with Gin
	c.router.Use(cors.Default())

	// Serve frontend static files
	//Always check how its build if its a package build
	c.router.Use(static.Serve("/", static.LocalFile("./frontend", true)))

	//Setup Ping
	c.router.GET("/ping", func(ctx *gin.Context) {
		type PingData struct {
			Message string `json:"message"`
			Time    string `json:"time"`
		}

		data := &PingData{
			Message: "Pong",
			Time:    time.Now().String(),
		}

		ctx.JSON(200, gin.H{
			"payload": data,
		})
	})

	if err := c.pool.RegisterReceiver(c); err != nil {
		c.log.Error(err)
	}

	// Setup route group for the API
	api := c.router.Group("/api")
	{
		api.GET("/data", func(ctx *gin.Context) {
			c.blackboard.Display.Time = time.Now().String()
			c.blackboard.Display.GoRoutineCount = runtime.NumGoroutine()
			ctx.JSON(http.StatusOK, gin.H{
				"payload": c.blackboard.Display,
			})
		})
		api.GET("/close", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"payload": "Server closing in 1 second",
			})
			go c.closeServer()
		})

		api.GET("/ws", func(context *gin.Context) {
			c.serveWs(c.pool, context.Writer, context.Request)
		})
	}

	// Start and run the server
	if err := c.router.Run(fmt.Sprintf(":%d", c.config.Data.Port)); err != nil {
		panic(err)
	}
}

func (c *RouterConstruct) closeServer() {
	time.Sleep(1 * time.Second)
	c.closeChan <- true
}

// define our WebSocket endpoint
func (c *RouterConstruct) serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	c.log.Info(r.Host)

	// upgrade this connection to a WebSocket
	// connection
	ws, err := websocket.Upgrade(w, r, c.log)
	if err != nil {
		c.log.Error(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	//go websocket.Writer(ws)
	//websocket.Reader(ws)

	client := &websocket.Client{
		Conn: ws,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func (c *RouterConstruct) FromClients(message websocket.WSMessage) bool {
	switch message.Msg.Type {
	case "data":
		go c.blackboard.UpdateDisplay()
		return true
	case "config":
		go c.config.UpdateConfig(message.Msg.Body)
		return true
	case "ping":
		go c.pool.BroadcastData("ping_pong", message.Msg)
		return true
	}
	return false
}
