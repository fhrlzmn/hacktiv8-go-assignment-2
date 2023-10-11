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
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type OrderHandlerImpl struct {
	service service.OrderService
}

func OrderHandlerInit(service service.OrderService) *OrderHandlerImpl {
	return &OrderHandlerImpl{service}
}

// Create Order godoc
// @Summary Create an order
// @Description Create an order with the input payload
// @Tags orders
// @Accept json
// @Produce json
// @Param order body dto.OrderRequest true "Order Payload"
// @Success 201 {object} entity.Order
// @Router /orders [post]
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

	response, err := oh.service.Create(order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errs.InternalServerError("Failed to create order"))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Get Order by ID godoc
// @Summary Get an order by id
// @Description Get an order by id
// @Tags orders
// @Accept json
// @Produce json
// @Param orderId path int true "Order ID"
// @Success 200 {object} entity.Order
// @Router /orders/{orderId} [get]
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

// Update Order godoc
// @Summary Update an order
// @Description Update an order with the input payload
// @Tags orders
// @Accept json
// @Produce json
// @Param orderId path int true "Order ID"
// @Param order body dto.OrderRequest true "Order Payload"
// @Success 200 {object} entity.Order
// @Router /orders/{orderId} [put]
func (oh *OrderHandlerImpl) Update(ctx *gin.Context) {
	orderId := ctx.Param("orderId")

	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errs.BadRequest("Invalid order ID"))
		return
	}

	var order entity.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, errs.BadRequest("Invalid request body"))
		return
	}

	response, err := oh.service.Update(orderIdInt, order)
	if err != nil {
		ctx.JSON(
			http.StatusNotFound,
			errs.NotFound("Order with ID "+orderId+" not found"),
		)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Delete Order godoc
// @Summary Delete an order
// @Description Delete an order by id
// @Tags orders
// @Accept json
// @Produce json
// @Param orderId path int true "Order ID"
// @Success 200 {object} string
// @Router /orders/{orderId} [delete]
func (oh *OrderHandlerImpl) Delete(ctx *gin.Context) {
	orderId := ctx.Param("orderId")

	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errs.BadRequest("Invalid order ID"))
		return
	}

	response, err := oh.service.Delete(orderIdInt)
	if err != nil {
		ctx.JSON(
			http.StatusNotFound,
			errs.NotFound("Order with ID "+orderId+" not found"),
		)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
