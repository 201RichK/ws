package main

import (
	"time"

	"github.com/gorilla/websocket"
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
