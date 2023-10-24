package global

type ResultCont struct {
	ErrCode int    `json:"errCode"` // 提示代码
	Message string `json:"message"` // 提示信息
}

func NewError(code int, msg string) ResultCont {
	return ResultCont{
		ErrCode: code,
		Message: msg,
	}
}
