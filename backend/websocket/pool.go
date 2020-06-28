package websocket

import (
	"backend/log"
	"errors"
)

type Receivers interface {
	FromClients(WSMessage)
}

type Pool struct {
	Register     chan *Client
	Unregister   chan *Client
	Clients      map[*Client]bool
	ToClients    chan Message
	FromClients  chan WSMessage
	AllReceivers []Receivers
}

func NewPool() *Pool {
	return &Pool{
		Register:     make(chan *Client),
		Unregister:   make(chan *Client),
		Clients:      make(map[*Client]bool),
		ToClients:    make(chan Message),
		FromClients:  make(chan WSMessage),
		AllReceivers: make([]Receivers, 0),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			log.Instance().Info("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				log.Instance().Info(client)
				writeJSON(client.Conn, Message{Type: "system", Body: "client connected"})
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			log.Instance().Info("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				writeJSON(client.Conn, Message{Type: "system", Body: "client disconnected"})
			}
			break
		case message := <-pool.ToClients:
			//log.Instance().Info("Sending message to all clients in Pool")
			for client, _ := range pool.Clients {
				writeJSON(client.Conn, message)
			}
			break
		case message := <-pool.FromClients:
			//log.Instance().Info("Sending message to all clients in Pool")
			for item := range pool.AllReceivers {
				go pool.AllReceivers[item].FromClients(message)
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
