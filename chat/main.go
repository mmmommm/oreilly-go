package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/mmmommm/go-sample/chat"
)

type templateHandler struct {
	filename string
	templ *template.Template
}

func (t *templateHandler) serveHTTP(w http.ResponseWriter, r *http.Request) {
	if t.templ == nil {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	}
	t.templ.Execute(w, nil)
}

func main() {
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	go r.run()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("listenandserve:", err);
	}
}