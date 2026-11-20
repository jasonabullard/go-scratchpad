package routes

import (
	"html/template"
	"net/http"

	"github.com/jasonabullard/go-scratchpad/todoapp/assets"
)

var tmpl = template.Must(template.ParseFS(assets.HTMLFiles, "*.html"))

func Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", index)
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServerFS(assets.StaticFiles)))
	registerTodos(mux)
}

func index(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index.html", todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
