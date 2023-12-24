package modelProduct

import (
	"errors"
	"time"
)

const (
	STATUS_ACTION  = 1
	STATUS_BLOCK   = -1
	STATUS_DELETED = -2
)

var (
	ErrBlocked      = errors.New("the product has been blocked")
	ErrDeleted      = errors.New("the product has been deleted")
	ErrNameIsBlank  = errors.New("name product cannot be blank")
	ErrPriceIsBlank = errors.New("price product cannot be blank")
)

type Product struct {
	Id         int        `json:"id" gorm:"column:id;"`
	Name       string     `json:"name" gorm:"column:name;"`
	Status     int        `json:"status" gorm:"column:status;"`
	Image      string     `json:"image" gorm:"column:image;"`
	Taste      string     `json:"taste" gorm:"column:taste;"`
	CategoryId int        `json:"category_id" gorm:"column:category_id;"`
	Price      int        `json:"price" gorm:"column:price;"`
	Discount   int        `json:"discount" gorm:"column:discount;"`
	CreatedAt  *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt  *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (Product) TableName() string {
	return "products"
}

type ProductCreate struct {
	Name  string `json:"name" gorm:"column:name;"`
	Price int    `json:"price" gorm:"column:price;"`
}

func (ProductCreate) TableName() string {
	return Product{}.TableName()
}
