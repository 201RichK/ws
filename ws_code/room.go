package main

//Room struct
type room struct {
	//Evry room has a name.
	name chan string

	//Register clients.
	client map[int]*Client

	//Hold the actual count
	count chan int

	//Uniqu Id
	index chan int
}

func newRoom() *room {
	return &room{
		name:   make(chan, string),
		client: make(map[int]*Client),
		count:  make(chan int),
		index:  make(chan int),
	}
}

func (r *room) Run() {
	for {
		select {}
	}
}
