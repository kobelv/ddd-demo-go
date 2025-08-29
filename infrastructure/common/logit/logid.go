package logit

import (
	"strconv"
	"sync/atomic"
	"time"
)

const LogIDKey = "logid"
const LogTraceID = "traceId"

var idx atomic.Int64

var now = func() time.Time {
	return time.Now()
}

// NewLogIDAny 获取一个新的logid
func NewLogIDAny() interface{} {
	usec := now().UnixNano() + idx.Add(1)
	logID := usec&0x7FFFFFFF | 0x80000000
	return strconv.FormatInt(logID, 10)
}
