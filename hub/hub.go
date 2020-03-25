package hub

import (
	"github.com/201RichK/ws/client"
)

//Hub maintains the set of active client and broadcasr msg to the client
type Hub struct {
	//Registered clients
	Clients map[*client.Client]bool

	//Inbound messages from the client
	Broadcast chan []byte

	//register request from the client
	Register chan *client.Client

	//Unregister request from the client
	Unregister chan *client.Client
}

//NewHub register client by adding the client pointer
//as a key in the client map
func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*client.Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *client.Client),
		Unregister: make(chan *client.Client),
	}
}

//Run methode
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
