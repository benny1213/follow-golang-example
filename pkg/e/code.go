package e

const (
	// SUCCESS : 请求成功 code
	SUCCESS = 200
	// ERROR : 请求错误 code
	ERROR = 500

	// InvalidParams : 参数错误 code
	InvalidParams = 400
	// ErrorExistTag : 已存在该标签名称
	ErrorExistTag = 10001
	// ErrorNotExistTag : 该标签不存在
	ErrorNotExistTag = 10002
	// ErrorNotExistArticle : 该文章不存在
	ErrorNotExistArticle = 10003

	// ErrorAuthCheckTokenFail : Token鉴权失败
	ErrorAuthCheckTokenFail = 20001
	// ErrorAuthCheckTokenTimeout : Token已超时
	ErrorAuthCheckTokenTimeout = 20002
	// ErrorAuthToken : Token生成失败
	ErrorAuthToken = 20003
	// ErrorAuth : Token错误
	ErrorAuth = 20004
)
