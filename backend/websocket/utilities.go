package websocket

import (
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	"io"
)

func writeJSON(conn *websocket.Conn, v interface{}) error {
	w, err := conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return err
	}
	err1 := jsoniter.ConfigFastest.NewEncoder(w).Encode(v)
	err2 := w.Close()
	if err1 != nil {
		return err1
	}
	return err2
}

func readJSON(conn *websocket.Conn, v interface{}) error {
	_, r, err := conn.NextReader()
	if err != nil {
		return err
	}
	err = jsoniter.ConfigFastest.NewDecoder(r).Decode(v)
	if err == io.EOF {
		// One value is expected in the message.
		err = io.ErrUnexpectedEOF
	}
	return err
}
