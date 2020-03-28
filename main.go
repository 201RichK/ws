package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

//single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Methode not allowed", http.StatusMethodNotAllowed)
		return
	}
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("views", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	hub := newHub()
	go hub.run()

	http.Handle("/", &templateHandler{filename: "index.html"})
	http.Handle("/ws", hub)

	var port string
	if os.Getenv("port") == "" {
		port = ":8080"
	}

	log.Fatal("ListenAndServe Error ::", http.ListenAndServe(port, nil))
}
