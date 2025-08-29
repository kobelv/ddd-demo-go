package dto

// OrderLog 日志记录相关标识
type OrderLog struct {
	LogID     string `json:"logId,omitempty"`
	TraceID   string `json:"traceId"`
	RequestID string `json:"requestId"`
}
