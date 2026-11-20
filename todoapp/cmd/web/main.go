package main

import (
	"log"
	"net/http"

	"github.com/jasonabullard/go-scratchpad/todoapp/routes"
)

func main() {
	mux := http.NewServeMux()
	routes.Register(mux)

	log.Fatal(http.ListenAndServe(":8090", mux))
}
