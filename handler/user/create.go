// API for creating account

package user

import (
	"accountserver/handler"
	"accountserver/pkg/errno"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func Create(c *gin.Context) {
	var r CreateRequest
	// Bind()检查Content-Type类型，将消息体按照指定的格式解析到struct中
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	pm := c.Param("username") // 读取并返回URL的参数值
	log.Infof("URL username: %s", pm)

	desc := c.Query("desc") // 读取并返回URL的地址参数
	log.Infof("URL key param desc: %s", desc)

	contentType := c.GetHeader("Content-Type") // 获取HTTP Header
	log.Infof("Header Content-Type: %s", contentType)

	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password) // 对内
	if r.Username == "" {
		handler.SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")), nil) // 对内
		return
	}
	if r.Password == "" {
		handler.SendResponse(c, fmt.Errorf("password is empty"), nil)
	}

	resp := CreateResponse{
		Username: r.Username,
	}
	handler.SendResponse(c, nil, resp)
}
