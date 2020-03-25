package main

import (
	"log"
	"net/http"
	"os"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Methode not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "tmpl/index.html")
}

func main() {
	hub := newHub()
	go hub.run()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	var port string
	if os.Getenv("port") == "" {
		port = ":8080"
	}

	log.Fatal("ListenAndServe Error ::", http.ListenAndServe(port, nil))
}
