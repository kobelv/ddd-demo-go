package dto

import (
	"ddd-demo-go/domain/entity"
)

type OrderConv struct {
}

func NewOrderConv() OrderConv {
	return OrderConv{}
}

// D2E 将dto转换为实体参数
func (s *OrderConv) D2E(dto *OrderDTO) *entity.OrderEntity {
	var en entity.OrderEntity

	en.OwnerID = dto.OwnerID
	en.Description = dto.Description

	var lines []entity.OrderLineEntity
	for _, l := range dto.Lines {
		var line entity.OrderLineEntity
		line.Price = l.Price
		line.ProductID = l.ProductID
		line.Quantity = l.Quantity
		line.Comment = l.Comment
		lines = append(lines, line)
	}

	return &en
}
