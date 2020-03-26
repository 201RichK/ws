package main

type room struct {
	//forwward is channel that hold incommind msg to other clients.
	forward chan []byte

	//join is a channel for client who want to join the room
	join chan *client

	//leave is channel for client wo want to leave the room
	leave chan *client

	//all current client in this room
	clients map[*client]bool
}
