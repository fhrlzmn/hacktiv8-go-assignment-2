package app

import (
	"fhrlzmn/hacktiv8-go/assignment-2/infrastructure/database"
	"fhrlzmn/hacktiv8-go/assignment-2/infrastructure/server"
)

func Start() {
	database.Init()
	gin := server.Init()
	gin.Run() // listen and serve on port 8080 if env PORT not set
}
