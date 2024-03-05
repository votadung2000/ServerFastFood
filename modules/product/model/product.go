package modelProduct

import (
	"errors"
	"fastFood/common"
	"strings"
)

const (
	STATUS_ACTION  = 1
	STATUS_BLOCK   = -1
	STATUS_DELETED = -2
)

const (
	EntityName = "Product"
)

var (
	ErrBlocked         = errors.New("the product has been blocked")
	ErrDeleted         = errors.New("the product has been deleted")
	ErrNameIsBlank     = errors.New("name product cannot be blank")
	ErrPriceIsBlank    = errors.New("price product cannot be blank")
	ErrCategoryIsBlank = errors.New("category product cannot be blank")
)

type Product struct {
	common.SQLModel
	Name        string  `json:"name" gorm:"column:name;"`
	ImageId     int     `json:"image_id" gorm:"column:image_id;"`
	Taste       string  `json:"taste" gorm:"column:taste;"`
	Price       float32 `json:"price" gorm:"column:price;"`
	CategoryId  int     `json:"category_id" gorm:"column:category_id;"`
	Discount    float32 `json:"discount" gorm:"column:discount;"`
	Status      int     `json:"status" gorm:"column:status;"`
	Description string  `json:"description" gorm:"column:description;"`
	Quantity    int     `json:"quantity" gorm:"column:quantity;"`
	Sold        int     `json:"sold" gorm:"column:sold;"`
	Featured    int     `json:"featured" gorm:"column:featured;"`
}

func (Product) TableName() string {
	return "products"
}

type ProductCreate struct {
	Name       string `json:"name" gorm:"column:name;"`
	Price      int    `json:"price" gorm:"column:price;"`
	CategoryId int    `json:"category_id" gorm:"column:category_id;"`
	ImageId    int    `json:"image_id" gorm:"column:image_id;"`
}

func (ProductCreate) TableName() string {
	return Product{}.TableName()
}

func (p *ProductCreate) Validate() error {
	p.Name = strings.TrimSpace(p.Name)

	if p.Name == "" {
		return ErrNameIsBlank
	}
	if p.Price == 0 {
		return ErrPriceIsBlank
	}
	if p.CategoryId == 0 {
		return ErrCategoryIsBlank
	}

	return nil
}

type ProductUpdate struct {
	Name     *string `json:"name" gorm:"column:name;"`
	Status   *int    `json:"status" gorm:"column:status;"`
	Image    *string `json:"image" gorm:"column:image;"`
	Taste    *string `json:"taste" gorm:"column:taste;"`
	Price    *int    `json:"price" gorm:"column:price;"`
	Discount *int    `json:"discount" gorm:"column:discount;"`
}

func (ProductUpdate) TableName() string {
	return Product{}.TableName()
}
