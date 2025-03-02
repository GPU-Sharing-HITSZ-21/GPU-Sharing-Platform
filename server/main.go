package main

import (
	"github.com/gin-gonic/gin"
	"gpu-sharing-platform/route"
	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()

	//cros配置
	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:35173","http://10.249.190.219:35173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length", "Authorization"},
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
