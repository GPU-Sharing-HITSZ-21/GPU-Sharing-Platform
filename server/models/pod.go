package models

import "gorm.io/gorm"

type Pod struct {
	gorm.Model
	Id       int
	PodName  string
	Username string
}
