package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeRouterInit(router *gin.Engine) {
	homeRouter := router.Group("home")
	{
		homeRouter.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "home",
			})
		})
	}
}
