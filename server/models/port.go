package models

import "gorm.io/gorm"

type Port struct {
	gorm.Model
	PortNum int `gorm:"column:port_num"`
	Status  int `gorm:"column:status"`
}
