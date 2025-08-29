package adapter

import (
	"github.com/gin-gonic/gin"

	"ddd-demo-go/domain/entity"
)

type UserAdapter interface {
	GetUserByID(ctx *gin.Context, id int64) (*entity.UserEntity, error)
}
