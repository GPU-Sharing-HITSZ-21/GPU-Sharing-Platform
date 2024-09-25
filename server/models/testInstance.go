package models

import "time"

type TestInstance struct {
	Id        int
	Name      string
	CreatedAt time.Time
}

func (testInstance TestInstance) TableName() string {
	return "test_instance"
}
