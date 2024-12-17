package main

import (
	"github.com/gin-gonic/gin"
	"gpu-sharing-platform/route"
)

func main() {
	router := gin.Default()

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
