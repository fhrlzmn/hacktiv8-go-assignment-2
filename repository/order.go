package repository

import (
	"fmt"

	"gorm.io/gorm"

	"fhrlzmn/hacktiv8-go/assignment-2/domain/entity"
)

type OrderRepository interface {
	Create(order entity.Order, items []entity.Item) (entity.Order, error)
}

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func OrderRepositoryInit(db *gorm.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{db}
}

func (or *OrderRepositoryImpl) Create(
	order entity.Order,
	items []entity.Item,
) (entity.Order, error) {
	err := or.db.Transaction(func(tx *gorm.DB) error {
		fmt.Println("ID", order.ID)

		for _, item := range items {
			fmt.Println("item.ID", item.ID)
		}

		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		return nil
	})

	return order, err
}
