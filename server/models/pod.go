package models

import "gorm.io/gorm"

type Pod struct {
	gorm.Model
	PodName    string
	Username   string
	SSHAddress string `gorm:"column:ssh_address"`
}
