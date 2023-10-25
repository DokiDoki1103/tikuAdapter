package errors

import "errors"

var (
	ErrDisable           = errors.New("关闭本接口的搜题")
	ErrTargetApiFlow     = errors.New("对方API接口限流")
	ErrTargetServerError = errors.New("对方服务器异常")
	ErrTargetNoAnswer    = errors.New("对方没有答案")
)
