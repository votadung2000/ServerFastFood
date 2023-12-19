package modelCategory

import (
	"errors"
	"fastFood/common"
)

const (
	STATUS_ACTION = 1
	STATUS_BLOCK  = -1
)

var (
	ErrBlocked = errors.New("the category has been blocked")
)

type Category struct {
	common.SQLModel
	Name   string `json:"name" gorm:"column:name;"`
	Status int    `json:"status" gorm:"column:status;"`
	Image  string `json:"image" gorm:"column:image;"`
}

func (Category) TableCategory() string {
	return "categories"
}

type CategoryUpdate struct {
	Name   string `json:"name" gorm:"column:name"`
	Status int    `json:"status" gorm:"column:status"`
	Image  string `json:"image" gorm:"column:image"`
}

func (CategoryUpdate) TableCategory() string {
	return Category{}.TableCategory()
}
