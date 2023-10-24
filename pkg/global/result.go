package global

type ResultCont struct {
	Code int    `json:"code"` //提示代码
	Msg  string `json:"msg"`  //提示信息
}

func NewError(code int, msg string) ResultCont {
	return ResultCont{
		Code: code,
		Msg:  msg,
	}
}
