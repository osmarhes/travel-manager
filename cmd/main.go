package main

import (
	"log"

	"github.com/osmarhes/travel-manager/cmd/server"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load("config/.env"); err != nil {
		log.Println(".env file not found. Continuing without it.")
	}
	server.Run()
}
