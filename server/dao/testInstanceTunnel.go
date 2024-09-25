package dao

import (
	"fmt"
	"gpu-sharing-platform/dao/dataSource"
	"gpu-sharing-platform/models"
)

var db = dataSource.DB

func SelectAllInstanceByPage(offset int, limit int) ([]models.TestInstance, error) {
	var instances []models.TestInstance

	// 执行分页查询
	if err := db.Offset(offset).Limit(limit).Find(&instances).Error; err != nil {
		return nil, err
	}
	fmt.Println(instances)
	return instances, nil
}
