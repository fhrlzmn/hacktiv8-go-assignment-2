package dto

type OrderRequest struct {
	OrderedAt    string        `json:"ordered_at"    binding:"required" example:"2021-01-01T00:00:00+07:00"`
	CustomerName string        `json:"customer_name" binding:"required"`
	Items        []ItemRequest `json:"items"         binding:"required"`
}

type ItemRequest struct {
	ItemCode    string `json:"item_code"   binding:"required"`
	Quantity    uint   `json:"quantity"    binding:"required"`
	Description string `json:"description" binding:"required"`
}
