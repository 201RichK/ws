/*************************************************************************
                     _____    _____    _____	_    _   _  __
 ____     ___	 _  |  __ \  |_   _|  /  ____| | |  | | | |/ /
|___ \	 / _ \  / | | |__) |   | |   |  |      | |__| | | ' /
  __) | | | | | | | |  _  /    | |   |  |      |  __  | |  <
 / __/	| |_| | | | | | \ \   _| |_  |  |____  | |  | | | . \
|_____|  \___/  |_| |_|  \_\ |_____|  \______| |_|  |_| |_|\_\ 201RichK powered this With 🥰❤️ and passion
***************************************************************************/
package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"
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
	http.Handle("/", &templateHandler{filename: "index.html"})
	http.Handle("/ws", hub)
	var port string = os.Getenv("port")
	if port == "" {
		port = ":8080"
	}

	log.Fatal("ListenAndServe Error ::> ", http.ListenAndServe(port, nil))
}
