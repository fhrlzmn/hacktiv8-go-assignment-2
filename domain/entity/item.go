package entity

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ItemCode    string `gorm:"not null;type:varchar(255)"`
	Quantity    uint   `gorm:"not null;type:int"`
	Description string `gorm:"not null;type:varchar(255)"`
	OrderID     uint   `gorm:"not null;type:int"`
}
