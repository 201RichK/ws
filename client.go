package main

import (
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

type (
	//Message struct
	message struct {
		//Message type
		msgType string

		//message body
		msg []byte
	}

	//Client struct
	Client struct {
		hub *hub

		//websocket connection for this client
		conn *websocket.Conn

		//messagge channel on witch msg are send
		send chan *message
	}
)

func newClient(conn *websocket.Conn) *Client {
	return &Client{
		conn: conn,
		send: make(chan *message),
	}
}

//ReadLoop Read meessage and pumps it to send channel
func (c *Client) ReadLoop() {
	defer close(c.send)
	for {
		_, p, err := c.conn.ReadMessage()
		if err != nil {
			log.Error("websocket read err ::> ", err)
			break
		}
		msg := message{msgType: "ex", msg: p}
		c.send <- &msg
	}
}

//WriteMessage write a message to the client
func (c *Client) WriteMessage(msg []byte) {
	err := c.conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Error("websocket write err ::> ", err)
	}
}
