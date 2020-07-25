package websocket

import (
	"backend/log"
	"errors"
	"github.com/ChicK00o/container"
	jsoniter "github.com/json-iterator/go"
)

type Receivers interface {
	FromClients(WSMessage) bool
}

type Pool struct {
	Register     chan *Client
	Unregister   chan *Client
	Clients      map[*Client]bool
	ToClients    chan Message
	FromClients  chan WSMessage
	AllReceivers []Receivers
	log          log.Logger
}

func init() {
	container.Singleton(func(logger log.Logger) *Pool {
		pool := &Pool{
			Register:     make(chan *Client),
			Unregister:   make(chan *Client),
			Clients:      make(map[*Client]bool),
			ToClients:    make(chan Message),
			FromClients:  make(chan WSMessage),
			AllReceivers: make([]Receivers, 0),
			log:          logger,
		}
		go pool.start()

		return pool
	})
}

func (pool *Pool) start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			pool.log.Info("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				pool.log.Info(client)
				_ = writeJSON(client.Conn, Message{Type: "system", Body: "client connected"}, pool.log)
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			pool.log.Info("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				_ = writeJSON(client.Conn, Message{Type: "system", Body: "client disconnected"}, pool.log)
			}
			break
		case message := <-pool.ToClients:
			//pool.log.Info("Sending message to all clients in Pool")
			for client, _ := range pool.Clients {
				_ = writeJSON(client.Conn, message, pool.log)
			}
			break
		case message := <-pool.FromClients:
			//pool.log.Info("Sending message to all clients in Pool")
			handled := false
			for item := range pool.AllReceivers {
				handled = handled || pool.AllReceivers[item].FromClients(message)
			}
			if !handled {
				data, _ := jsoniter.ConfigFastest.MarshalToString(message.Msg.Body)
				pool.log.Error("Unhandled message type : ", message.Msg.Type, " with data : ", data)
			}
		}
	}
}

func (pool *Pool) RegisterReceiver(receiver Receivers) error {
	for item := range pool.AllReceivers {
		if pool.AllReceivers[item] == receiver {
			return errors.New("Item already Exists")
		}
	}
	pool.AllReceivers = append(pool.AllReceivers, receiver)
	return nil
}

func (pool *Pool) UnRegisterReceiver(receiver Receivers) error {
	index := -1
	for item := range pool.AllReceivers {
		if pool.AllReceivers[item] == receiver {
			index = item
			break
		}
	}
	if index < 0 {
		return errors.New("Item not found")
	}
	pool.AllReceivers = append(pool.AllReceivers[:index], pool.AllReceivers[index+1:]...)
	return nil
}

func (pool *Pool) BroadcastText(message string) {
	pool.ToClients <- Message{Type: "generic", Body: message}
}

func (pool *Pool) BroadcastData(messageType string, data interface{}) {
	pool.ToClients <- Message{Type: messageType, Body: data}
}
