package service

import (
	"fmt"
	"net/http"

	"fhrlzmn/hacktiv8-go/assignment-2/domain/dto"
	"fhrlzmn/hacktiv8-go/assignment-2/domain/entity"
	"fhrlzmn/hacktiv8-go/assignment-2/repository"
)

type OrderService interface {
	Create(order entity.Order) (*dto.ApiResponse[any], error)
	GetById(orderId int) (*dto.ApiResponse[any], error)
	Update(orderId int, order entity.Order) (*dto.ApiResponse[any], error)
	Delete(orderId int) (*dto.ApiResponse[any], error)
}

type OrderServiceImpl struct {
	orderRepository repository.OrderRepository
}

func OrderServiceInit(orderRepository repository.OrderRepository) *OrderServiceImpl {
	return &OrderServiceImpl{orderRepository}
}

func (os *OrderServiceImpl) Create(
	order entity.Order,
) (*dto.ApiResponse[any], error) {
	newOrder, err := os.orderRepository.Create(order)
	if err != nil {
		return nil, err
	}

	response := dto.ApiResponse[any]{
		StatusCode: http.StatusCreated,
		Message:    "Order created successfully",
		Data:       newOrder,
	}

	return &response, nil
}

func (os *OrderServiceImpl) GetById(orderId int) (*dto.ApiResponse[any], error) {
	order, err := os.orderRepository.GetById(orderId)
	if err != nil {
		return nil, err
	}

	response := dto.ApiResponse[any]{
		StatusCode: http.StatusOK,
		Message:    "Order retrieved successfully",
		Data:       order,
	}

	return &response, nil
}

func (os *OrderServiceImpl) Update(orderId int, order entity.Order) (*dto.ApiResponse[any], error) {
	order, err := os.orderRepository.Update(orderId, order)
	if err != nil {
		return nil, err
	}

	response := dto.ApiResponse[any]{
		StatusCode: http.StatusOK,
		Message:    "Order updated successfully",
		Data:       order,
	}

	return &response, nil
}

func (os *OrderServiceImpl) Delete(orderId int) (*dto.ApiResponse[any], error) {
	err := os.orderRepository.Delete(orderId)
	if err != nil {
		return nil, err
	}

	response := dto.ApiResponse[any]{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Order with ID %d deleted successfully", orderId),
		Data:       nil,
	}

	return &response, nil
}
