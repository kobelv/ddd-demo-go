package repository

import (
	"github.com/gin-gonic/gin"

	"ddd-demo-go/domain/entity"
	"ddd-demo-go/domain/repository"
	"ddd-demo-go/infrastructure/common/db"
	"ddd-demo-go/infrastructure/po"
)

type OrderRepositoryImpl struct {
	db          *db.DB
	orderPOConv *po.OrderPOConv
}

func NewOrderRepository(db *db.DB, conv *po.OrderPOConv) repository.OrderRepository {
	return &OrderRepositoryImpl{db: db, orderPOConv: conv}
}

func (repo *OrderRepositoryImpl) Save(ctx *gin.Context, en *entity.OrderEntity) (string, error) {
	return "1", nil

	// todo:
	// 下面代码需连接正确的mysql服务后才能启用
	// if repo.db.DB == nil {
	// 	return "", nil
	// }

	// orderPo := repo.orderPOConv.E2P(en)

	// res := repo.db.WithContext(ctx).Create(&orderPo)
	// if res.Error != nil {
	// 	return "", res.Error
	// }

	// return orderPo.ID, nil
}
