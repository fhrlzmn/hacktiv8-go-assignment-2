package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "fhrlzmn/hacktiv8-go/assignment-2/docs"
	"fhrlzmn/hacktiv8-go/assignment-2/handler"
	"fhrlzmn/hacktiv8-go/assignment-2/infrastructure/database"
	"fhrlzmn/hacktiv8-go/assignment-2/repository"
	"fhrlzmn/hacktiv8-go/assignment-2/service"
)

// @title Hacktiv8 Go Assignment 2 API Documentation
// @version 1.0.0
// @description This is a sample server for Hacktiv8 Go Assignment 2 API Documentation.
// @termsOfService https:/swagger.io/terms/
// @contact.name Fahrul Zaman
// @contact.email fhrlzmn26@gmail.com
// @license.name MIT
// @license.url http://mit-license
// @host localhost:8080
// @BasePath /api
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

	api := r.Group("/api")
	{
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		order := api.Group("/orders")
		{
			order.POST("/", orderHandler.Create)
			order.GET("/", orderHandler.GetAll)
			order.GET("/:orderId", orderHandler.GetById)
			order.PUT("/:orderId", orderHandler.Update)
			order.DELETE("/:orderId", orderHandler.Delete)
		}

	}

	return r
}
