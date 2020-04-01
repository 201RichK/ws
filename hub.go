package main

import (
	"net/http"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type hub struct {
	rooms map[string]*room
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

//Create a new hub
func newHub() *hub {
	return &hub{
		rooms: make(map[string]*room),
	}
}

//If room does"nt exist create it and return it
func (h *hub) getRoom(name string) *room {
	if _, ok := h.rooms[name]; !ok {
		h.rooms[name] = newRoom(&name)
	}
	return h.rooms[name]
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
	go room.run()
	id := room.join(newClient(conn))

	// time.Sleep(time.Second) //sleep to get Id ////REVIEWS LATER
	log.Info(room.clients, "id === ", id)

	//read from the client send channel and broadcast it
	go room.HandleMsg(id)

	//read from client and if this loop break then client disconnected
	room.clients[id].ReadLoop()
	room.leave <- id
}
