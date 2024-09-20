package main

import (
	"gpu-sharing-platform/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})

	routes.HomeRouterInit(router)

	router.Run(":1024")
}
