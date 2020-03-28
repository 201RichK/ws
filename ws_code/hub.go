package main

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

//Hub maintains the set of active client and broadcasr msg to the client
type hub struct {
	//hub for control room
	hub map[string]*room

	//Registered clients
	clients map[*Client]bool

	//Inbound messages from the client
	broadcast chan message

	//register request from the client
	register chan *Client

	//Unregister request from the client
	unregister chan *Client
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

//NewHub register client by adding the client pointer
//as a key in the client map
func newHub() *hub {
	return &hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

//Run methode
func (h *hub) run() {
	for {
		select {
		case client := <-h.register:
			//joining
			log.Info(client)
			h.clients[client] = true
		case client := <-h.unregister:
			//leaving
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			//send msg to all client
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
