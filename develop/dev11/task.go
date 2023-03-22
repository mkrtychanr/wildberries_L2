package main

import (
	"log"
)

func main() {
	server, err := newServer()
	if err != nil {
		log.Fatal(err)
	}
	server.SetupHandlers()
	server.Up()
}
