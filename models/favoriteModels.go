package models

import "time"

type Favorite struct {
	Id        int        `json:"id" gorm:"column:id;"`
	UserId    int        `json:"user_id" gorm:"column:user_id;"`
	ProductId int        `json:"product_id" gorm:"column:product_id;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (Favorite) TableFavorites() string {
	return "favorites"
}
