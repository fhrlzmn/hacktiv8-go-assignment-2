package main

import (
	"log"

	"github.com/joho/godotenv"

	"fhrlzmn/hacktiv8-go/assignment-2/infrastructure/database"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	database.Init()
}
