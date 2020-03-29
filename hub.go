package main

import (
	"net/http"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type hub struct {
	hub map[string]*room
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	id int
)

//Create a new hub
func newHub() *hub {
	return &hub{
		hub: make(map[string]*room),
	}
}

//If room does"nt exist create it and return it
func (h *hub) getRoom(name string) *room {
	if _, ok := h.hub[name]; !ok {
		h.hub[name] = newRoom(&name)
	}
	return h.hub[name]
}

//ServeHTTP method allow us to get ws conn
func (h *hub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("ServeHTTP upgrade err ::> ", err)
		return
	}
	defer conn.Close()
	room := h.getRoom("generale")
	log.Info(room)
	go room.run()
	room.join <- conn
	go room.HandleMsg(id)
	room.clients[id].ReadLoop()
}
