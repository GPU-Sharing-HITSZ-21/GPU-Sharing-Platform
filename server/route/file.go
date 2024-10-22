package route

import (
	"github.com/gin-gonic/gin"
	"gpu-sharing-platform/service/file"
)

func FileRouterInit(router *gin.Engine) {
	fileRouter := router.Group("/file")
	{
		fileRouter.POST("/upload", file.HandleFileUpload)
	}
}
