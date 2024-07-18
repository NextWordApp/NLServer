package err

// 定义一个映射，将状态码映射到对应的消息字符串
var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
}

// GetMsg 根据状态码获取对应的消息字符串
func GetMsg(code int) string {
	// 从映射中获取消息字符串
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	// 如果状态码不存在，返回默认的错误消息
	return MsgFlags[ERROR]
}
