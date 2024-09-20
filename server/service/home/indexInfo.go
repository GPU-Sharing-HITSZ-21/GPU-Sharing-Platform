package home

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gpu-sharing-platform/dao"
)

func GetIndexInfo(context *gin.Context) {
	// 设置中国时区
	location, _ := time.LoadLocation("Asia/Shanghai")
	currentTime := time.Now().In(location).Format("2006-01-02 15:04:05")

	context.JSON(http.StatusOK, gin.H{
		"dateTime": currentTime,
	})
}

func GetTestInstance(context *gin.Context) {
	testInstances, err := dao.SelectAllInstanceByPage(0, 10)
	if err == nil {
		context.JSON(http.StatusOK, testInstances)
	} else {
		context.JSON(http.StatusOK, err)
	}
}
