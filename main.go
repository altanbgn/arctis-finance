package main

import (
	"log"

	"github.com/altanbgn/arctis-finance/server"
)

func main() {
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
