package websocket

import (
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type string      `json:"type"`
	Body interface{} `json:"body"`
}

type WSMessage struct {
	Type int     `json:"type"`
	Msg  Message `json:"msg"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		_ = c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			c.Pool.log.Error(err)
			continue
		}
		data := Message{}
		err = jsoniter.ConfigFastest.Unmarshal(p, &data)
		if err != nil {
			c.Pool.log.Error(err)
			continue
		}
		c.Pool.FromClients <- WSMessage{Type: messageType, Msg:data}
		c.Pool.log.Info("Message Received: ", string(p))
	}
}
