package websocket

import (
	"backend/log"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type        int         `json:"type"`
	MessageType string      `json:"message_type"`
	Body        interface{} `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Instance().Error(err)
			return
		}
		c.Pool.FromClients <- Message{Type: messageType, Body: string(p)}
		log.Instance().Info("Message Received: ", string(p))
	}
}
