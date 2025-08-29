package po

import "ddd-demo-go/domain/entity"

type OrderPOConv struct{}

func NewOrderPOConv() *OrderPOConv {
	return &OrderPOConv{}
}

func (conv *OrderPOConv) E2P(en *entity.OrderEntity) (po *OrderPO) {
	return nil
}
