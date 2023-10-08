package entity

type Item struct {
	Base
	ItemCode    string `gorm:"not null;type:varchar(255)" json:"item_code"`
	Quantity    uint   `gorm:"not null;type:int"          json:"quantity"`
	Description string `gorm:"not null;type:varchar(255)" json:"description"`
	OrderID     uint   `gorm:"not null;type:int"          json:"-"`
}
