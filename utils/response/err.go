package response

// 自定义业务错误状态码
type ErrorCode int

const (
	SettingsError ErrorCode = 1001 //  系统错误
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "系统错误",
	}
)
