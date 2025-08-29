package appservice

import (
	"github.com/gin-gonic/gin"

	"ddd-demo-go/application/dto"
	"ddd-demo-go/domain/adapter"
	"ddd-demo-go/domain/entity"
	"ddd-demo-go/domain/repository"
)

type OrderAS struct {
	dtoConv   dto.OrderConv
	userAdp   adapter.UserAdapter
	orderRepo repository.OrderRepository
}

func NewOrderAS(dtoConv dto.OrderConv, userAdp adapter.UserAdapter, orderRepo repository.OrderRepository) *OrderAS {
	return &OrderAS{
		dtoConv:   dtoConv,
		userAdp:   userAdp,
		orderRepo: orderRepo,
	}
}

// PlaceOrder 进入下单应用程序，编排业务逻辑
func (app *OrderAS) PlaceOrder(ctx *gin.Context, dto *dto.OrderDTO) (*entity.OrderEntity, error) {
	// dto参数校验
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	orderEn := app.dtoConv.D2E(dto)

	// 示例：rpc调用远程账号系统
	// _, err := app.userAdp.GetUserByID(ctx, dto.OwnerID)
	// if err != nil {
	// 	return nil, err
	// }

	orderEn.Create()

	id, err := app.orderRepo.Save(ctx, orderEn)
	if err != nil {
		return nil, err
	}
	orderEn.ID = id

	return orderEn, nil
}
