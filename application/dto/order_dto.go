package dto

type OrderDTO struct {
	ID          string         `json:"id"` // 订单id
	Description string         `json:"description"`
	OwnerID     int64          `json:"ownerId"`
	Lines       []OrderLineDTO `json:"lines"`
}

type OrderLineDTO struct {
	ID        string  `json:"id"` // 订单行id
	ProductID string  `json:"productID"`
	Quantity  float64 `json:"quantity"`
	Price     float64 `json:"price"`
	Comment   string  `json:"comment"`
}

// Validate 参数检验
func (dto *OrderDTO) Validate() error {
	return nil
}
