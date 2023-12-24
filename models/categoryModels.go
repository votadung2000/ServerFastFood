package models

import (
	"fastFood/common"
)

type Categories struct {
	common.SQLModel
	Name   string `json:"name" gorm:"column:name;"`
	Status int    `json:"status" gorm:"column:status;"` // Status: 1 - action, 2 - block
	Image  string `json:"image" gorm:"column:image;"`
}

func (Categories) TableCategory() string {
	return "categories"
}
