package entity

import (
	"time"
)

type Order struct {
	Base
	CustomerName string    `gorm:"not null;type:varchar(255)"  json:"customer_name"`
	OrderedAt    time.Time `                                   json:"ordered_at"`
	Items        []Item    `gorm:"constraint:OnUpdate:CASCADE" json:"items"`
}
