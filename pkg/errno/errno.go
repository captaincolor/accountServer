// 错误处理

package errno

import (
	"fmt"
)

type Errno struct {
	Code    int
	Message string
}

// Errno实现了error接口
func (err Errno) Error() string {
	return err.Message
}

type Err struct {
	Code    int
	Message string
	Err     error
}

func New(errno *Errno, err error) *Err {
	return &Err{Code: errno.Code, Message: errno.Message, Err: err}
}

func (err *Err) Add(message string) error {
	err.Message += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

// Err实现error接口
func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}

func IsInternalServerError(err error) bool {
	code, _ := DecodeErr(err)
	return code == InternalServerError.Code
}

func IsErrBind(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrBind.Code
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typ := err.(type) {
	case *Errno:
		return typ.Code, typ.Message
	case *Err:
		return typ.Code, typ.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
