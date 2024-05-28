package modelCategory

import (
	"fastFood/common"
	modelProduct "fastFood/modules/product/model"
	"strings"
)

const (
	STATUS_ACTION  = 1
	STATUS_BLOCK   = -1
	STATUS_DELETED = -2
)

const (
	EntityName = "Category"
)

var (
	ErrBlocked     = "the category has been blocked"
	ErrNameIsBlank = "the name category cannot be blank"
)

type Category struct {
	common.SQLModel
	Name     string                  `json:"name" gorm:"column:name;"`
	Status   int                     `json:"status" gorm:"column:status;"`
	ImageId  int                     `json:"-" gorm:"column:image_id;"`
	Image    *common.PreloadImage    `json:"image" gorm:"foreignKey:ImageId"`
	Products []*modelProduct.Product `json:"products" gorm:"foreignKey:CategoryId"`
}

func (Category) TableName() string {
	return "categories"
}

type CategoryUpdate struct {
	Name    *string `json:"name" gorm:"column:name"`
	Status  *int    `json:"status" gorm:"column:status"`
	ImageId *int    `json:"image_id" gorm:"column:image_id;"`
}

func (CategoryUpdate) TableName() string {
	return Category{}.TableName()
}

type CategoryCreate struct {
	Name    string `json:"name" gorm:"column:name"`
	ImageId int    `json:"image_id" gorm:"column:image_id;"`
}

func (CategoryCreate) TableName() string {
	return Category{}.TableName()
}

func (i *CategoryCreate) Validate() error {
	i.Name = strings.TrimSpace(i.Name)

	if i.Name == "" {
		return ErrValidateRequest(ErrNameIsBlank, "ERR_NAME_IS_BLANK")
	}

	return nil
}
