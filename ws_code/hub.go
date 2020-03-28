package main

import (
	"net/http"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

//Hub coontrol a bunch of room
type Hub struct {
	hub map[string]*room
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

//if room exist return it else create it
func (h *Hub) getroom(name string) *room {
	if _, ok := h.hub[name]; !ok {
		h.hub[name] = newRoom(room)
	}
	return h.hub[name]
}

func (h *Hub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("upgrade error ::", err)
	}

	defer c.Close()
}
