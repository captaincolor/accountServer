// 统一自定义错误码, code唯一标识错误，message向前端展示错误信息
// 1开头，系统级错误
// 2开头，普通错误

package errno

var (
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct"} // 没有传入任何参数时触发
	ErrUserNotFound     = &Errno{Code: 20102, Message: "The user was not found"}
)
