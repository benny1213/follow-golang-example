package e

// MsgFlag : 错误码flag释义
var MsgFlag = map[int]string{
	SUCCESS:                    "ok",
	ERROR:                      "fail",
	InvalidParams:              "请求参数错误",
	ErrorExistTag:              "已存在该标签名称",
	ErrorNotExistTag:           "该标签不存在",
	ErrorNotExistArticle:       "该文章不存在",
	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",
}

// GetMsg : 获取错误信息
func GetMsg(code int) string {
	msg, ok := MsgFlag[code]
	if ok {
		return msg
	}
	return MsgFlag[ERROR]
}
