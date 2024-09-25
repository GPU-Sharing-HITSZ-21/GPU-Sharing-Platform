package route

import (
	"gpu-sharing-platform/service/home"
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

		homeRouter.GET("/time", home.GetIndexInfo)

		homeRouter.GET("/get_test_instance", home.GetTestInstance)
	}
}
