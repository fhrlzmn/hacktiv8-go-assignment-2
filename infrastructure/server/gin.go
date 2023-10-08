package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fhrlzmn/hacktiv8-go/assignment-2/handler"
	"fhrlzmn/hacktiv8-go/assignment-2/infrastructure/database"
	"fhrlzmn/hacktiv8-go/assignment-2/repository"
	"fhrlzmn/hacktiv8-go/assignment-2/service"
)

func Init() *gin.Engine {
	r := gin.Default()           // router
	db := database.GetInstance() // db

	orderRepository := repository.OrderRepositoryInit(db)
	orderService := service.OrderServiceInit(orderRepository)
	orderHandler := handler.OrderHandlerInit(orderService)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "active",
			"message": "Welcome to the API!",
		})
	})

	r.Group("/")
	{
		order := r.Group("/orders")
		order.POST("/", orderHandler.Create)
	}

	return r
}
