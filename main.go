package main

import (
	"log"

	"github.com/joho/godotenv"

	"fhrlzmn/hacktiv8-go/assignment-2/cmd/app"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	app.Start()
}
