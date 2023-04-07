package main

import (
	_ "go-training/internal/migrations"
	"go-training/internal/server"
)

func main() {
	port := 3001

	server.Start(port)
}
