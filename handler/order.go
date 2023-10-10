package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"fhrlzmn/hacktiv8-go/assignment-2/domain/entity"
	"fhrlzmn/hacktiv8-go/assignment-2/pkg/errs"
	"fhrlzmn/hacktiv8-go/assignment-2/service"
)

type OrderHandler interface {
	Create(ctx *gin.Context)
	GetById(ctx *gin.Context)
}

type OrderHandlerImpl struct {
	service service.OrderService
}

func OrderHandlerInit(service service.OrderService) *OrderHandlerImpl {
	return &OrderHandlerImpl{service}
}

func (oh *OrderHandlerImpl) Create(ctx *gin.Context) {
	order := entity.Order{}
	items := []entity.Item{}

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, errs.BadRequest("Invalid request body"))
		return
	}

	if order.CustomerName == "" {
		ctx.JSON(http.StatusBadRequest, errs.BadRequest("Customer name cannot be empty"))
		return
	}

	items = order.Items // insert the items from request body
	if items == nil || len(items) == 0 {
		ctx.JSON(http.StatusBadRequest, errs.BadRequest("Items cannot be empty"))
		return
	}

	response, err := oh.service.Create(order, items)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errs.InternalServerError("Failed to create order"))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (oh *OrderHandlerImpl) GetById(ctx *gin.Context) {
	orderId := ctx.Param("orderId")

	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errs.BadRequest("Invalid order ID"))
		return
	}

	response, err := oh.service.GetById(orderIdInt)
	if err != nil {
		ctx.JSON(
			http.StatusNotFound,
			errs.NotFound("Order with ID "+orderId+" not found"),
		)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
