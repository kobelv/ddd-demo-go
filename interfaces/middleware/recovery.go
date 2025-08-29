package middleware

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"

	"ddd-demo-go/infrastructure/common/logit"
)

func Recovery(logger logit.LoggerInterface, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			logit.AddAllLevel(c, logit.Any("status", http.StatusInternalServerError))
			if err := recover(); err != any(nil) {
				var brokenPipe bool
				if ne, ok := any(err).(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				logid, _ := c.Get("logid")
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c, c.Request.URL.Path,
						logit.Any("logid", logid),
						logit.Any("error", err),
						logit.Any("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(any(err).(error)) // nolint: errcheck
					c.Abort()
					return
				}
				if stack {
					logger.Error(c, "recovery from panic",
						logit.Any("logid", logid),
						logit.Any("error", err),
						logit.Any("request", string(httpRequest)),
						logit.Any("stack", string(debug.Stack())),
					)
				} else {
					logger.Error(c, "recovery from panic",
						logit.Any("logid", logid),
						logit.Any("error", err),
						logit.Any("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
