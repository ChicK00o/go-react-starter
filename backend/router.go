package main

import (
	"backend/blackboard"
	"backend/config"
	"backend/log"
	"backend/websocket"
	"fmt"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"os"
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
}

func NewRouterConstruct(l log.Logger, u *Utilities, b *blackboard.Blackboard, c *config.Config) *RouterConstruct {
	gin.SetMode(gin.ReleaseMode)
	return &RouterConstruct{log: l, utilities: u, router: gin.Default(), blackboard: b, config:c}
}

func (c *RouterConstruct) startRouter(portNumber int) {
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

	c.pool = websocket.NewPool()
	go c.pool.Start()
	if err := c.pool.RegisterReceiver(c); err != nil {
		c.log.Error(err)
	}

	// Setup route group for the API
	api := c.router.Group("/api")
	{
		api.GET("/data", func(ctx *gin.Context) {
			c.blackboard.DataHolder.Time = time.Now().String()
			ctx.JSON(http.StatusOK, gin.H{
				"payload": c.blackboard.DataHolder,
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

	go c.listenForBlackboard()
	// Start and run the server
	if err := c.router.Run(fmt.Sprintf(":%d", portNumber)); err != nil {
		panic(err)
	}
}

func (c *RouterConstruct) closeServer() {
	time.Sleep(1 * time.Second)
	os.Exit(3)
}

// define our WebSocket endpoint
func (c *RouterConstruct) serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	c.log.Info(r.Host)

	// upgrade this connection to a WebSocket
	// connection
	ws, err := websocket.Upgrade(w, r)
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

func (c *RouterConstruct) FromClients(message websocket.WSMessage) {
	switch message.Msg.Type {
	case "data":
		c.blackboard.UpdateDisplay()
		break
	case "config":
		c.config.UpdateConfig(message.Msg.Body)
		break
	default:
		data, _ := jsoniter.ConfigFastest.MarshalToString(message.Msg.Body)
		c.log.Error("Unhandled message type : ", message.Msg.Type, " with data : ", data)
		go c.pool.BroadcastData("ping_pong", message.Msg)
	}
}

func (c *RouterConstruct) listenForBlackboard() {
	c.blackboard.ListenerAttached = true
	for {
		_ = <-c.blackboard.UpdateChannel
		c.blackboard.DataHolder.Time = time.Now().String()
		c.blackboard.DataHolder.GoRoutineCount = runtime.NumGoroutine()
		go c.pool.BroadcastData("display", c.blackboard.DataHolder)
		go c.pool.BroadcastData("config", c.config.Data)
	}
}
