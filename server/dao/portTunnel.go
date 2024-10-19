package dao

import (
	"errors"
	"gorm.io/gorm"
	"gpu-sharing-platform/models"
)

// ClaimPort 尝试占用指定的端口
func ClaimPort() (int, error) {
	// 从数据库中获取当前最大端口号
	currentMaxPort, err := GetCurrentMaxPort()
	if err != nil {
		return -1, err // 处理错误
	}

	// 尝试找到一个未被占用的端口
	for portNum := currentMaxPort + 1; portNum <= 65535; portNum++ { // 假设最大端口为 65535
		if !CheckPort(portNum) {
			// 插入端口到数据库
			port := &models.Port{
				PortNum: portNum,
			}
			_, err := InsertPort(port)
			if err != nil {
				return -1, err
			}
			return portNum, nil // 返回成功的端口号
		}
	}

	return -1, errors.New("no available ports")
}

// CheckPort 检查指定的端口是否已被占用
func CheckPort(portNum int) bool {
	var count int64
	result := db.Model(&models.Port{}).Where("port_num = ?", portNum).Count(&count)

	if result.Error != nil {
		// 处理错误
	}

	if count > 0 {
		return true
	} else {
		return false
	}
}

// GetPortByNum 根据端口号查找端口
func GetPortByNum(portNum int) (*models.Port, error) {
	var port models.Port
	result := db.Where("port_num = ?", portNum).First(&port)
	if result.Error != nil {
		return nil, result.Error
	}
	return &port, nil
}

// InsertPort 将端口插入数据库
func InsertPort(port *models.Port) (int, error) {
	result := db.Create(port)
	if result.Error != nil {
		return -1, result.Error
	}
	return int(port.ID), nil
}

// GetCurrentMaxPort 获取当前最大端口号
func GetCurrentMaxPort() (int, error) {
	var port models.Port
	result := db.Order("port_num DESC").First(&port) // 获取最大端口号
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 30000, nil // 如果没有记录，返回初始port
		}
		return -1, result.Error
	}
	return port.PortNum, nil
}
