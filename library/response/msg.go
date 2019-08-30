package response

var MsgFlags = map[int]string{
	SUCCESS:                  "ok",
	ERROR:                    "fail",
	INVALID_PARAMS:           "请求参数错误",
	ERROR_EXIST:              "已存在资源",
	ERROR_NOT_EXIST:          "资源不存在",
	ERROR_AUTH_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN_CREATE:  "Token生成失败",
	ERROR_AUTH:               "Token错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
