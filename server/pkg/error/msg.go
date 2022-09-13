package e

var MsgFlags = map[int]string{
	SUCCESS:                 "OK",
	INTERNAL_ERROR:          "内部服务器错误",
	INVALID_PARAMS:          "请求参数错误",
	ERROR_EXIST_TAG:         "标签已存在",
	ERROR_NOT_EXIST_ARTICLE: "文章不存在",
	ERROR_NOT_EXIST_TAG:     "标签不存在",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "token不存在",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "token已过期",
	ERROR_AUTH_TOKEN:               "token生成失败",
	ERROR_AUTH:                     "token错误",

	ERROR_USER: "无效的用户",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[INTERNAL_ERROR]
}
