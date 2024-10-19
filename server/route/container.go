package route

import (
	"github.com/gin-gonic/gin"
	"gpu-sharing-platform/service/k8sHandler"
)

func ContainerRouterInit(router *gin.Engine) {
	containerRouter := router.Group("container")
	{
		containerRouter.GET("/test", k8sHandler.CreateTestPod)
		containerRouter.GET("/terminal", k8sHandler.HandleExecWebSocket)
		containerRouter.POST("/create", k8sHandler.CreatePodByUser)
		containerRouter.POST("/delete", k8sHandler.DeletePodByName)
		containerRouter.GET("/myPods", k8sHandler.GetPodByUser)

	}
}
