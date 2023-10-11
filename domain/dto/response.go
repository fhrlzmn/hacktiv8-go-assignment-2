package dto

import (
	"time"

	"fhrlzmn/hacktiv8-go/assignment-2/domain/entity"
)

type ApiResponse[T any] struct {
	StatusCode uint   `json:"status_code"`
	Message    string `json:"message"`
	Data       T      `json:"data"`
}

type OrderResponse struct {
	ID           uint           `json:"id"`
	OrderedAt    time.Time      `json:"ordered_at"`
	CustomerName string         `json:"customer_name"`
	Items        []ItemResponse `json:"items"`
}

type ItemResponse struct {
	ID          uint   `json:"id"`
	ItemCode    string `json:"item_code"`
	Quantity    uint   `json:"quantity"`
	Description string `json:"description"`
}

func GetOrderResponse(order entity.Order) OrderResponse {
	var orderResponse OrderResponse
	var itemsResponse []ItemResponse

	for _, item := range order.Items {
		itemsResponse = append(itemsResponse, ItemResponse{
			ID:          item.ID,
			ItemCode:    item.ItemCode,
			Quantity:    item.Quantity,
			Description: item.Description,
		})
	}

	orderResponse = OrderResponse{
		ID:           order.ID,
		OrderedAt:    order.OrderedAt,
		CustomerName: order.CustomerName,
		Items:        itemsResponse,
	}

	return orderResponse
}

func GetDetailedOrderResponse(order entity.Order) entity.Order {
	return order
}
