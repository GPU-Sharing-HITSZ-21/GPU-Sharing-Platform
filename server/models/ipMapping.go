package models

import (
	"gorm.io/gorm"
)

// IpMapping 公网到内网 IP 的映射
type IpMapping struct {
	gorm.Model
	PublicIP  string `gorm:"column:public_ip;not null"`  // 公网 IP
	PrivateIP string `gorm:"column:private_ip;not null"` // 内网 IP
}

func (ipMapping IpMapping) TableName() string {
	return "ip_mapping"
}
