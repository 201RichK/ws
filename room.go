package main

import (
	log "github.com/sirupsen/logrus"
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
	leave     chan int
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
		leave:     make(chan int),
		broadcast: make(chan []byte),
		sendTo:    make(chan *msgTo),
		sendEx:    make(chan *msgTo),
	}
}

func (r *room) join(client *Client) int {
	r.index++
	r.clients[r.index] = client
	r.count++
	return r.index
}

func (r *room) run() {
	for {
		select {
		case id := <-r.leave:
			//Remove client from "room"
			r.count--
			delete(r.clients, id)
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

//handle message
func (r *room) HandleMsg(id int) {
	for {
		if r.clients[id] == nil {
			log.Info("client not exist")
			break
		}
		send := <-r.clients[id].send
		if send.msgType == "ex" {
			r.sendEx <- &msgTo{id, send.msg}
		} else {
			r.broadcast <- send.msg
		}
	}
}
