package global

const (
	// FileNotFound 文件未找到
	FileNotFound = iota + 5000001

	// FileParseError 文件解析错误
	FileParseError

	// FileHashError  计算文件hash失败
	FileHashError

	// FileUploadError 文件上传失败
	FileUploadError
)

var (
	// ErrorParam 错误的参数
	ErrorParam = NewError(400400, "错误的参数")

	// ErrorParseFile 解析文件出错
	ErrorParseFile = NewError(FileParseError, "解析文件出错")

	// ErrorFileNotFound 文件未找到
	ErrorFileNotFound = NewError(FileNotFound, "文件未找到")

	// ErrorFileHashError 计算文件hash失败
	ErrorFileHashError = NewError(FileHashError, "计算文件hash失败")

	// ErrorFileUploadError 文件上传失败
	ErrorFileUploadError = NewError(FileUploadError, "文件上传失败")
)
