package e

var MsgFlags = map[int]string{
	SUCCESS:                    "ok",
	ERROR:                      "fail",
	InvalidParams:              "请求参数错误",
	HaveSignup:                 "已注册",
	ErrorActivityTimeout:       "活动过期",
	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token过期",
	ErrorAuthToken:             "Token生成超时",
	ErrorAuth:                  "Token失败",
	ErrorNotCompare:            "不匹配",
	ErrorDatabase:              "数据库操作出错，请重试",
	ErrorAuthNotFound:          "Token不能为空",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
