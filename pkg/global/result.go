package global

// ResultCont 返回错误结果
type ResultCont struct {
	ErrCode int    `json:"errCode"` // 提示代码
	Message string `json:"message"` // 提示信息
}

// NewError 创建错误
func NewError(code int, msg string) ResultCont {
	return ResultCont{
		ErrCode: code,
		Message: msg,
	}
}
