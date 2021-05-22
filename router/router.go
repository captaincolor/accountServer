// 绑定http路径和处理函数
package router

import (
	"accountServer/handler/sc"
	"accountServer/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 加载路由
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 设置http header
	g.Use(gin.Recovery()) // ginRecovery()恢复api server，避免bug或异常情况导致下一次请求调用被影响
	g.Use(middleware.NoCache) // 强制browser不使用缓存
	g.Use(middleware.Options) // browser跨域请求设置
	g.Use(middleware.Secure) // 安全设置
	g.Use(mw...)
	// 404 handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Incorrect API route.")
	})

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