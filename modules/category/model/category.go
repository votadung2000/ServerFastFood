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
	ErrBlocked     = errors.New("the category has been blocked")
	ErrNameIsBlank = errors.New("name category cannot be blank")
)

type Category struct {
	common.SQLModel
	Name   string `json:"name" gorm:"column:name;"`
	Status int    `json:"status" gorm:"column:status;"`
	Image  string `json:"image" gorm:"column:image;"`
}

func (Category) TableName() string {
	return "categories"
}

type CategoryUpdate struct {
	Name   string `json:"name" gorm:"column:name"`
	Status int    `json:"status" gorm:"column:status"`
	Image  string `json:"image" gorm:"column:image"`
}

func (CategoryUpdate) TableName() string {
	return Category{}.TableName()
}

type CategoryCreate struct {
	Name string `json:"name" gorm:"column:name"`
}

func (CategoryCreate) TableName() string {
	return Category{}.TableName()
}
