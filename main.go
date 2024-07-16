package main

import (
	"github.com/shekhar-patil/gobank/app"
	"log"
)

func main() {
	store, err := app.NewPostgresStore()

	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := app.NewAPIServer(":3000", *store)
	server.Run()
}
