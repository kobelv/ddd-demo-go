package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ddd-demo-go/infrastructure/common/cache"
	"ddd-demo-go/infrastructure/common/response"
)

// 心跳健康请求构造结构体
type Health struct {
	response.HTTPResponseInterface
	redis *cache.Redis
}

func NewHealth(
	resp response.HTTPResponseInterface,
	redis *cache.Redis,
) *Health {
	return &Health{HTTPResponseInterface: resp, redis: redis}
}

func (h *Health) Liveness(c *gin.Context) {
	h.RenderJSONResponse(c, http.StatusOK, nil, nil)
	return
}

func (h *Health) Readiness(c *gin.Context) {
	if err := h.redis.Ping().Err(); err != nil {
		h.RenderJSONResponse(c, http.StatusInternalServerError, nil, err)
		return
	}
	h.RenderJSONResponse(c, http.StatusOK, nil, nil)
	return
}
