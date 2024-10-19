package dao

import "gpu-sharing-platform/models"

// GetPublicIpByPrivateIp 根据内网 IP 查找对应的公网 IP
func GetPublicIpByPrivateIp(privateIp string) (string, error) {
	var mapping models.IpMapping
	result := db.Where("private_ip = ?", privateIp).First(&mapping)
	if result.Error != nil {
		return "", result.Error
	}
	return mapping.PublicIP, nil
}
