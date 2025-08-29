package repository

import (
	"github.com/gin-gonic/gin"

	"ddd-demo-go/domain/entity"
)

type OrderRepository interface {
	Save(ctx *gin.Context, en *entity.OrderEntity) (string, error)
}
