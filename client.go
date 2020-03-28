package main

import (
	"bytes"
	"time"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
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
}

//readPump pumps message from the websocket conneection to the hub.
func (c *client) readPump() {
	defer func() {
		c.hub.unregister <- c
		defer c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Errorf("Error :: %v", err)
			}
			break
		}
		msg = bytes.TrimSpace(bytes.Replace(msg, newline, space, -1))
		c.hub.broadcast <- msg
	}
}

//writePump pumps message from the hub to the websocket connection
func (c *client) writePump() {

}
