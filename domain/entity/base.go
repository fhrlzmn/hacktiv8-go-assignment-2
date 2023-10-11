package entity

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	CreatedAt time.Time      `                                json:"created_at,omitempty"`
	UpdatedAt time.Time      `                                json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index"                    json:"-"`
}
