package main

import (
	"log"

	"github.com/your-username/your-project/server"
)

func main() {
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}