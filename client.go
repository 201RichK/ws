package main

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
type client struct {
	hub *hub

	//websocket connection for this client
	conn *websocket.Conn

	//Buffered channel on witch msg are send
	send chan []byte

	//room is the room this client is chating
	room *room
}

//readPump pumps message from the websocket conneection to the hub.
func (c *client) readPump() {

}

//writePump pumps message from the hub to the websocket connection
func (c *client) writePump() {

}

//ServeWs handle websockets request
func serveWs(h *hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Error("serveWs upgrade err :: ", err)
		return
	}

	client := &client{
		hub:  h,
		conn: conn,
		send: make(chan []byte, 256),
	}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}
