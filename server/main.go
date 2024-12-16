package main

import (
	"github.com/gin-gonic/gin"
	"gpu-sharing-platform/route"
)

func main() {
	router := gin.Default()

	//cros配置
	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, // 允许的源，可以根据需要添加多个源
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, // 允许的 HTTP 方法
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 允许的请求头
        ExposeHeaders:    []string{"Content-Length", "Authorization"}, // 允许客户端获取的响应头
        AllowCredentials: true,
    }))

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})

	route.HomeRouterInit(router)
	route.ContainerRouterInit(router)
	route.UserRouterInit(router)
	route.FileRouterInit(router)
	route.JobRouterInit(router)

	err := router.Run(":31024")
	if err != nil {
		// todo 告警
		return
	}
}
