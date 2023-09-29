package entity

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerName string `gorm:"not null;type:varchar(255)"`
	OrderedAt    time.Time
	Items        []Item `gorm:"constraint:OnUpdate:CASCADE"`
}
