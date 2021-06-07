// 统一自定义错误码

package errno

var (
	OK = &Errno{Code: 0, Message: "OK"}

	// 1开头，系统级错误
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	// 2开头，普通错误
	ErrUserNotFound = &Errno{Code: 20102, Message: "The user was not found."}
)
