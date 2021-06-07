// api for creating account

package user

import (
	"accountserver/pkg/errno"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"net/http"
)

func Create(c *gin.Context) {
	var r struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var err error
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		return
	}

	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password) // 对内
	if r.Username == "" {
		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")) // 对内
		log.Errorf(err, "Error!")                                                                       // 对内
	}

	if errno.IsErrUserNotFound(err) {
		log.Debug("err type is ErrUserNotFound") // 对内
	}

	if r.Password == "" {
		err = fmt.Errorf("password is empty") // 对外
	}

	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
