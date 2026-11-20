package routes

import (
	"encoding/json/v2"
	"io"
	"net/http"
	"slices"
	"uuid"

	"github.com/jasonabullard/go-scratchpad/todoapp/response"
)

type Todo struct {
	ID   string `json:"id"`
	Body string `json:"body"`
	Done bool   `json:"done"`
}

var todos = []Todo{}

func list(w http.ResponseWriter, r *http.Request) {
	response.OK(w, todos)
}

func get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	for _, todo := range todos {
		if todo.ID == id {
			response.OK(w, todo)
			return
		}
	}

	response.Fail(w, http.StatusNotFound, "Todo not found")
}

func create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Fail(w, http.StatusInternalServerError, "Unable to read body")
	}

	var todo Todo
	if err = json.Unmarshal(body, &todo); err != nil {
		response.Fail(w, http.StatusInternalServerError, "Cannot unmarshal data")
	}

	todo.ID = uuid.New().String()
	todos = append(todos, todo)

	response.OK(w, todo)
}

func delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	for i, todo := range todos {
		if todo.ID == id {
			todos = slices.Delete(todos, i, i+1)
			response.OK(w, todo)
			return
		}
	}

	response.Fail(w, http.StatusNotFound, "ID not found")
}

func registerTodos(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/todos", list)
	mux.HandleFunc("GET /api/v1/todos/{id}", get)
	mux.HandleFunc("POST /api/v1/todos", create)
	mux.HandleFunc("DELETE /api/v1/todos/{id}", delete)
}
