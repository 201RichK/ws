package client

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

const (
	//Time allowed to write msg to the peer.
	writeWait = 10 * time.Second

	//Time allowed to read the next pong msg from the peer.
	pongWait = 60 * time.Second

	//send pings to peer with this period. Must be less than pongwait
	pingPeriod = (pongWait * 9) / 10

	//Maximum message size allowed from peer
	maxMessageSize = 522
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}

	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

//Client is a middleman between the websocket connection and the hub.
type Client struct {
	Hub *hub.Hub

	//websocket connection
	Conn *websocket.Conn

	//Buffered channel of outbound meesages.
	Send chan []byte
}

//readPump pumps message from the websocket conneection to the hub.
func (c *Client) readPump() {

}

//writePump pumps message from the hub to the websocket connection
func (c *Client) writePump() {

}

//ServeWs handle websockets request
func ServeWs(h *hub.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Error("serveWs upgrade err :: ", err)
		return
	}

	client := &Client{
		Hub:  h,
		Conn: conn,
		Send: make(chan []byte, 256),
	}
	client.Hub.Register <- client

	go client.writePump()
	go client.readPump()
}
