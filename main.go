package main

import (
	"log"

	"github.com/gimmy/awesome-api/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalf("server exited with error: %v", err)
	}
}
