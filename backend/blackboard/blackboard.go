package blackboard

import (
	"backend/config"
	"backend/log"
	"backend/websocket"
	"github.com/ChicK00o/container"
	"runtime"
	"time"
)

type Blackboard struct {
	logger        log.Logger
	Display       *DisplayDataHolder
	UpdateChannel chan bool
	InternalData  *CustomInternalData
	pool          *websocket.Pool
	con           *config.Config
}

type DisplayDataHolder struct {
	Message        string             `json:"message"`
	Time           string             `json:"time"`
	GoRoutineCount int                `json:"go_routine_count"`
	Data           *CustomDisplayData `json:"data"`
}

func init() {
	container.Singleton(func(logger log.Logger, p *websocket.Pool, c *config.Config) *Blackboard {
		dataHolder := &DisplayDataHolder{
			Data: InitialCustomData(),
		}
		board := &Blackboard{
			logger:        logger,
			Display:       dataHolder,
			UpdateChannel: make(chan bool),
			InternalData:  InitialCustomInternalData(),
			pool:          p,
			con:           c,
		}
		go board.listenForUpdate()
		return board
	})
}

func (b *Blackboard) UpdateDisplay() {
	b.UpdateChannel <- true
}

func (b *Blackboard) listenForUpdate() {
	for {
		_ = <-b.UpdateChannel
		b.Display.Time = time.Now().String()
		b.Display.GoRoutineCount = runtime.NumGoroutine()
		go b.pool.BroadcastData("display", b.Display)
		go b.pool.BroadcastData("config", b.con.Data)
	}
}
