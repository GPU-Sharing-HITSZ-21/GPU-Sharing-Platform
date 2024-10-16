package route

import (
	"github.com/gin-gonic/gin"
	"gpu-sharing-platform/service/userControl"
)

func UserRouterInit(router *gin.Engine) {
	userRouter := router.Group("user")
	{
		// 用户注册
		userRouter.POST("/register", userControl.RegisterUser)

		// 用户登录
		userRouter.POST("/login", userControl.LoginUser)

		//// 获取用户信息
		//userRouter.GET("/:id", userControl.GetUser)
		//
		//// 更新用户信息
		//userRouter.PUT("/:id", userControl.UpdateUser)
		//
		//// 删除用户
		//userRouter.DELETE("/:id", userControl.DeleteUser)
	}
}
