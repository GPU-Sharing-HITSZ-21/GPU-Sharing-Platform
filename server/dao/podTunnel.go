package dao

import (
	"gpu-sharing-platform/models"
)

func InsertPod(pod *models.Pod) (int, error) {
	result := db.Create(pod)
	if result.Error != nil {
		return -1, result.Error
	} else {
		return int(pod.ID), nil
	}
}

func GetPodById(podId int) (*models.Pod, error) {
	var pod models.Pod

	// 使用 First 方法根据 podId 查找 Pod
	result := db.First(&pod, podId) // 查找主键等于 podId 的记录

	if result.Error != nil {
		return nil, result.Error
	}

	return &pod, nil
}

func GetPodsByUsername(username string) ([]models.Pod, error) {
	var pods []models.Pod

	// 根据 UserID 查找对应的 Pods
	result := db.Where("user_name = ?", username).Find(&pods)
	if result.Error != nil {
		return nil, result.Error
	}

	return pods, nil
}
