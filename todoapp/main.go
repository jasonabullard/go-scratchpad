package main

import (
	"log"
	"net/http"

	"github.com/jasonabullard/go-scratchpad/todoapp/routes"
)

type Todo struct {
	ID        string
	Body      string
	Completed bool
}

func main() {
	mux := http.NewServeMux()
	routes.Register(mux)

	log.Fatal(http.ListenAndServe(":8090", mux))
}
