package models

import "gorm.io/gorm"

type Pod struct {
	gorm.Model
	PodName    string `gorm:"column:pod_name"`                       // 自定义列名
	Username   string `gorm:"column:username"`                       // 自定义列名
	SSHAddress string `gorm:"column:ssh_address" json:"ssh_address"` // 自定义列名
	PortNum    int    `gorm:"column:port_num" json:"port_num"`
}
