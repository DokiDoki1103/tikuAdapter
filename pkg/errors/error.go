package errors

import "errors"

var (
	// ErrTargetAPIFlow 对方API接口限流
	ErrTargetAPIFlow = errors.New("对方API接口限流")
	// ErrTargetServerError 对方服务器异常
	ErrTargetServerError = errors.New("对方服务器异常")
	// ErrTargetNoAnswer 对方没有答案
	ErrTargetNoAnswer = errors.New("对方没有答案")
	// ErrParserJSON 解析json错误
	ErrParserJSON = errors.New("解析json错误")
	// ErrTokenRequired 需要提供卡密
	ErrTokenRequired = errors.New("需要提供卡密")
)
