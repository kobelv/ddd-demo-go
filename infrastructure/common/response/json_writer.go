package response

import (
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"

	"ddd-demo-go/infrastructure/common/errors"
	"ddd-demo-go/infrastructure/common/logit"
)

type HTTPResponseInterface interface {
	RenderJSONResponse(ctx *gin.Context, statusCode int, data any, err error)
}

// 返回响应内容结构体
type body struct {
	Code      int64  `json:"code"`
	Msg       string `json:"msg"`
	Data      any    `json:"data"`
	Timestamp int64  `json:"timestamp"`
	LogID     string `json:"logid"`
}

type responseWriter struct {
	logger logit.LoggerInterface
}

func NewHTTPResponseWriter(logger logit.LoggerInterface) HTTPResponseInterface {
	return &responseWriter{logger: logger}
}

func (h *responseWriter) RenderJSONResponse(ctx *gin.Context, statusCode int, data any, err error) {
	logID, _ := ctx.Get("logid")
	logIDString := logID.(string)
	response := &body{
		Code:      int64(errors.Success),
		Msg:       "success",
		Data:      data,
		Timestamp: time.Now().Unix(),
		LogID:     logIDString,
	}
	if err != nil {
		response.Code = int64(errors.GetType(err))
		response.Msg = errors.Cause(err).Error()
	}

	if isNil(data) {
		response.Data = struct{}{}
	}
	h.logger.Info(ctx, "us响应ui结果", logit.Any("response", response))
	if statusCode == 0 {
		statusCode = http.StatusOK
	}
	ctx.JSON(statusCode, response)
}

func isNil(v any) bool {
	valueOf := reflect.ValueOf(v)
	k := valueOf.Kind()
	switch k {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return valueOf.IsNil()
	default:
		return v == nil
	}
}
