// 绑定http路径和处理函数

package router

import (
	"accountserver/handler/sc"
	"accountserver/handler/user"
	"accountserver/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

// loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// middlewares, set http headers
	g.Use(gin.Recovery()) // 恢复svr，避免bug或异常情况导致下一次请求调用被影响
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 handler
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Incorrect API route.")
	})

	// routes for creating account
	u := g.Group("/v1/user")
	{
		u.POST("", user.Create)
	}

	// sc(serverCheck)分组主要用来检查API服务器的状态
	// 在sc分组下注册了多个http路径，分别路由到sc中的处理函数
	svrCheck := g.Group("/sc")
	{
		svrCheck.GET("/health", sc.HealthCheck)
		svrCheck.GET("/disk", sc.DiskCheck)
		svrCheck.GET("/cpu", sc.CPUCheck)
		svrCheck.GET("/ram", sc.RAMCheck)
	}

	return g
}
