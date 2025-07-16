package main

import (
	"log"

	"travel-manager/cmd/server"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env file not found. Continuing without it.")
	}
	server.Run()
}
