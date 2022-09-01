package errcode

var (
	Success                           = NewError(0, "成功")
	OK                                = NewError(0, "成功")
	ServerError                       = NewError(1000, "服务内部错误")
	InvalidParams                     = NewError(1001, "入参错误")
	NotFound                          = NewError(1002, "查询资源不存在")
	UnauthorizedAuthNotExist          = NewError(1003, "鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError            = NewError(1004, "鉴权失败，Token错误")
	UnauthorizedTokenTimeout          = NewError(1005, "鉴权失败，Token超时")
	UnauthorizedTokenGenerate         = NewError(1006, "鉴权失败，Token生成失败")
	UnauthorizedTokenSignatureInvalid = NewError(1007, "鉴权失败，Token签名异常")
	TooManyRequests                   = NewError(1008, "请求过多")
)
