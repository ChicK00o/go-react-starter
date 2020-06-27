package websocket

import (
	"backend/log"
	"github.com/gorilla/websocket"
	"net/http"
)

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// We'll need to check the origin of our connection
	// this will allow us to make requests from our React
	// development server to here.
	// For now, we'll do no checking and just allow any connection
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Instance().Error(err)
		return ws, err
	}
	return ws, nil
}

//// define a reader which will listen for
//// new messages being sent to our WebSocket
//// endpoint
//func Reader(conn *websocket.Conn) {
//	for {
//		// read in a message
//		messageType, p, err := conn.ReadMessage()
//		if err != nil {
//			log.Instance().Error(err)
//			return
//		}
//		// print out that message for clarity
//
//		log.Instance().Info(string(p))
//
//		type dataType struct {
//			Message string `json:"message"`
//			Time    string `json:"time"`
//		}
//
//		data := &dataType{
//			Message: "received : " + string(p),
//			Time:    time.Now().String(),
//		}
//
//		toSend, err := jsoniter.ConfigFastest.MarshalToString(&data)
//		if err != nil {
//			log.Instance().Error(err)
//			return
//		}
//
//		if err := conn.WriteMessage(messageType, []byte(toSend)); err != nil {
//			log.Instance().Error(err)
//			return
//		}
//	}
//}
//
//func Writer(conn *websocket.Conn) {
//	for {
//		log.Instance().Debug("Sending")
//		messageType, r, err := conn.NextReader()
//		if err != nil {
//			log.Instance().Error(err)
//			return
//		}
//		w, err := conn.NextWriter(messageType)
//		if err != nil {
//			log.Instance().Error(err)
//			return
//		}
//		if _, err := io.Copy(w, r); err != nil {
//			log.Instance().Error(err)
//			return
//		}
//		if err := w.Close(); err != nil {
//			log.Instance().Error(err)
//			return
//		}
//	}
//}
