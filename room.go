package main

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type msgTo struct {
	userID  int
	message []byte
}

type room struct {
	name      *string
	clients   map[int]*Client
	count     int
	index     int
	join      chan *websocket.Conn
	leave     chan *int
	broadcast chan []byte
	sendTo    chan *msgTo
	sendEx    chan *msgTo
}

func newRoom(name *string) *room {
	return &room{
		name:      name,
		clients:   make(map[int]*Client),
		count:     0,
		index:     0,
		join:      make(chan *websocket.Conn),
		leave:     make(chan *int),
		broadcast: make(chan []byte),
		sendTo:    make(chan *msgTo),
		sendEx:    make(chan *msgTo),
	}
}

func (r *room) run() {
	for {
		select {
		case conn := <-r.join:
			//add a conn to client map
			r.index++
			r.clients[r.index] = newClient(conn)
			r.count++
			id = r.index
		case id := <-r.leave:
			//Remove client from room
			r.count--
			delete(r.clients, *id)
		case msgto := <-r.sendTo:
			//Send to specifiq client
			r.clients[msgto.userID].WriteMessage(msgto.message)
		case msg := <-r.broadcast:
			//Broadcast to evry client
			for _, cl := range r.clients {
				cl.WriteMessage(msg)
			}
		case msgEx := <-r.sendEx:
			//Broadcast to all except
			for id, client := range r.clients {
				if id != msgEx.userID {
					client.WriteMessage(msgEx.message)
				}
			}
		}
	}
}

/* Handle messages */
func (r *room) HandleMsg(id int) {
	for {
		if r.clients[id] == nil {
			break
		}
		send := <-r.clients[id].send
		logrus.Info("msg :: ", send.msg)
		if send.msgType == "ex" {
			r.sendEx <- &msgTo{id, send.msg}
		} else {
			r.broadcast <- send.msg
			logrus.Info("msg :: ", send.msg)
		}
	}
}
