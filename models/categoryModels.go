package models

import "time"

type Categories struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Name      string     `json:"name" gorm:"column:name;"`
	Status    int        `json:"status" gorm:"column:status;"`
	Image     string     `json:"image" gorm:"column:image;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (Categories) TableCategory() string {
	return "categories"
}
