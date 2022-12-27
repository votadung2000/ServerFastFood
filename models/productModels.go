package models

import "time"

type Products struct {
	Id         int        `json:"id" gorm:"column:id;"`
	Name       string     `json:"name" gorm:"column:name;"`
	Status     int        `json:"status" gorm:"column:status;"` // Status: 1 - action, 2 - block
	Image      string     `json:"image" gorm:"column:image;"`
	Taste      string     `json:"taste" gorm:"column:taste;"`
	CategoryId int        `json:"category_id" gorm:"column:category_id;"`
	Price      int        `json:"price" gorm:"column:price;"`
	Discount   int        `json:"discount" gorm:"column:discount;"`
	CreatedAt  *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt  *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (Products) TableProducts() string {
	return "products"
}
