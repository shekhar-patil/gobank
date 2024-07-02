package main

import (
	"github.com/shekhar-patil/gobank/app"
)

func main() {
	server := app.NewAPIServer(":3000")
	server.Run()
}
