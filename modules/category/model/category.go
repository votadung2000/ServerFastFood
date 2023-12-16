package modelCategory

import (
	"fastFood/common"
)

const (
	STATUS_ACTION = 1
	STATUS_BLOCK  = -1
)

type Category struct {
	common.SQLModel
	Name   string `json:"name" gorm:"column:name;"`
	Status int    `json:"status" gorm:"column:status;"` // Status: 1 - action, -1 - block
	Image  string `json:"image" gorm:"column:image;"`
}

func (Category) TableCategory() string {
	return "categories"
}
