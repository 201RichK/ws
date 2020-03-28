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
		//websocket connection for this client
		conn *websocket.Conn

		//messagge channel on witch msg are send
		send chan message
	}
)

//ReadLoop Read meessage and pumps it to send channel
func (c *Client) ReadLoop() {}

//WriteMessage write a message to the client
func (c *Client) WriteMessage() {}
