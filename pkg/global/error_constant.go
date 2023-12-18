package global

var (
	// ErrorParam 错误的参数
	ErrorParam = NewError(400400, "错误的参数")

	// ErrorParseFile 解析文件出错
	ErrorParseFile = NewError(400500, "解析文件出错")
)
