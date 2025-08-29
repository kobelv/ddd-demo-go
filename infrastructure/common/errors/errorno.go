package errors

const (
	Success                      = ErrorType(0)    // 请求成功
	SignErr                      = ErrorType(1000) // sign错误
	FailedToDecodeRequestBodyErr = ErrorType(1001) // 请求体解析失败
	ParamsInvalidErr             = ErrorType(1002) // 参数不合法

	RPCResponseError = ErrorType(1003)

	DBError      = ErrorType(1004) // 读取数据库错误
	NoType       = ErrorType(1005) // 未知错误类型
	SystemError  = ErrorType(1006) // 系统错误
	RedisErr     = ErrorType(1007) // redis错误
	LogConfigErr = ErrorType(1011) // 日志配置格式错误
	AppConfigErr = ErrorType(1012) // app配置格式错误
)
