package entity

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"-"`
	CreatedAt time.Time      `                                json:"-"`
	UpdatedAt time.Time      `                                json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"                    json:"-"`
}
