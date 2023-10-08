package service

import (
	"net/http"

	"fhrlzmn/hacktiv8-go/assignment-2/domain/dto"
	"fhrlzmn/hacktiv8-go/assignment-2/domain/entity"
	"fhrlzmn/hacktiv8-go/assignment-2/repository"
)

type OrderService interface {
	Create(order entity.Order, items []entity.Item) (*dto.ApiResponse[any], error)
}

type OrderServiceImpl struct {
	orderRepository repository.OrderRepository
}

func OrderServiceInit(orderRepository repository.OrderRepository) *OrderServiceImpl {
	return &OrderServiceImpl{orderRepository}
}

func (os *OrderServiceImpl) Create(
	order entity.Order,
	items []entity.Item,
) (*dto.ApiResponse[any], error) {
	_, err := os.orderRepository.Create(order, items)
	if err != nil {
		return nil, err
	}

	response := dto.ApiResponse[any]{
		StatusCode: http.StatusCreated,
		Message:    "Order created successfully",
		Data:       order,
	}

	return &response, nil
}
