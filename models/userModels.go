package models

import "time"

type Users struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Name      string     `json:"name" gorm:"column:name;"`
	UserName  string     `json:"username" gorm:"column:username;"`
	PassWord  string     `json:"password" gorm:"column:password;"`
	Status    int        `json:"status" gorm:"column:status;"` // Status: 1 - action, 2 - block
	Image     string     `json:"image" gorm:"column:image;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (Users) TableUsers() string {
	return "users"
}