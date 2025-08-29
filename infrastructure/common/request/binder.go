package request

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"ddd-demo-go/infrastructure/common/logit"
)

type BindInterface interface {
	Bind(*gin.Context, any) error
}

// 请求日志追踪结构体
type request struct {
	logger logit.LoggerInterface
}

func NewRequest(logger logit.LoggerInterface) BindInterface {
	return &request{logger: logger}
}

func (r *request) Bind(req *gin.Context, v any) error {
	b := binding.Default(req.Request.Method, req.ContentType())
	if err := req.ShouldBindWith(v, b); err != nil {
		return err
	}
	return nil
}
