package route

import (
	"github.com/gin-gonic/gin"
	"gpu-sharing-platform/service/k8sHandler"
)

func ContainerRouterInit(router *gin.Engine) {
	containerRouter := router.Group("container")
	{
		containerRouter.GET("test", k8sHandler.CreateTestPod)
	}
}
