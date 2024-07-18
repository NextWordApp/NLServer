package err

// 定义常量，用于表示不同的状态码
const (
	SUCCESS        = 200 // 成功
	ERROR          = 500 // 失败
	INVALID_PARAMS = 400 // 请求参数错误

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001 // Token鉴权失败
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002 // Token已超时
	ERROR_AUTH_TOKEN               = 20003 // Token生成失败
	ERROR_AUTH                     = 20004 // Token错误
)
