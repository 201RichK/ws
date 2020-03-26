package main

//Hub maintains the set of active client and broadcasr msg to the client
type hub struct {
	//Registered clients
	clients map[*client]bool

	//Inbound messages from the client
	broadcast chan []byte

	//register request from the client
	register chan *client

	//Unregister request from the client
	unregister chan *client
}

//NewHub register client by adding the client pointer
//as a key in the client map
func newHub() *hub {
	return &hub{
		clients:    make(map[*client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *client),
		unregister: make(chan *client),
	}
}

//Run methode
func (h *hub) run() {
	for {
		select {
		case client := <-h.register:
			//joining
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
