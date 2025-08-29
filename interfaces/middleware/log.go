package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"ddd-demo-go/application/dto"
	"ddd-demo-go/infrastructure/common/logit"
	"ddd-demo-go/infrastructure/common/response"
)

func Logger(logger logit.LoggerInterface, response response.HTTPResponseInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		logID := logit.NewLogIDAny()
		c.Set(logit.LogIDKey, logID)
		cost := time.Since(start)
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			response.RenderJSONResponse(c, http.StatusOK, nil, err)
			return
		}
		uiLog := dto.OrderLog{}
		decode := json.NewDecoder(bytes.NewReader(body))
		decode.UseNumber()
		err = decode.Decode(&uiLog)
		if err == nil {
			c.Set(logit.LogIDKey, uiLog.LogID)
			c.Set(logit.LogTraceID, uiLog.TraceID)
		}
		logID, _ = c.Get(logit.LogIDKey)
		logger.Info(c, "ui请求入参", logit.Any("请求入参:", body))
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		logit.AddAllLevel(
			c,
			logit.Any("status", c.Writer.Status()),
			logit.Any("method", c.Request.Method),
			logit.Any("remote_addr", c.ClientIP()),
			logit.Any("path", path),
			logit.Any("logid", logID),
			logit.Any("query", c.Request.URL.RawQuery),
			logit.Any("useragent", c.Request.UserAgent()),
			logit.Any("cost", cost),
		)
		c.Next()
	}
}
