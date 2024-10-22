package route

import (
	"github.com/gin-gonic/gin"
	"gpu-sharing-platform/service/k8sHandler"
)

func JobRouterInit(router *gin.Engine) {
	jobRouter := router.Group("job")
	{
		jobRouter.POST("/start", k8sHandler.StartTrainingJob)
	}

}
