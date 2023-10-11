package repository

import (
	"gorm.io/gorm"

	"fhrlzmn/hacktiv8-go/assignment-2/domain/entity"
)

type OrderRepository interface {
	Create(order entity.Order) (entity.Order, error)
	GetById(orderId int) (entity.Order, error)
	Update(orderId int, updatedOrder entity.Order) (entity.Order, error)
	Delete(orderId int) error
}

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func OrderRepositoryInit(db *gorm.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{db}
}

func (or *OrderRepositoryImpl) Create(order entity.Order) (entity.Order, error) {
	err := or.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		return nil
	})

	return order, err
}

func (or *OrderRepositoryImpl) GetById(orderId int) (entity.Order, error) {
	var order entity.Order

	err := or.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload("Items").First(&order, orderId).Error; err != nil {
			return err
		}

		return nil
	})

	return order, err
}

func (or *OrderRepositoryImpl) Update(
	orderId int,
	updatedOrder entity.Order,
) (entity.Order, error) {
	var order entity.Order

	err := or.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload("Items").First(&order, orderId).Error; err != nil {
			return err
		}

		if err := tx.Model(&order).Updates(&updatedOrder).Error; err != nil {
			return err
		}

		for _, item := range updatedOrder.Items {
			var existingItem entity.Item
			tx.Where("order_id = ? AND item_code = ?", order.ID, item.ItemCode).
				First(&existingItem)

			if existingItem.ID != 0 {
				tx.Model(&existingItem).Updates(&item)
				continue
			}

			item.OrderID = uint(order.ID)
			tx.Create(&item)
		}

		return nil
	})

	return order, err
}

func (or *OrderRepositoryImpl) Delete(orderId int) error {
	var order entity.Order

	err := or.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&order, orderId).Error; err != nil {
			return err
		}

		if err := tx.Select("Items").Delete(&order).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}
