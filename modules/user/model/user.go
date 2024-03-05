package modelUser

import "fastFood/common"

const (
	STATUS_ACTION  = 1
	STATUS_BLOCK   = -1
	STATUS_DELETED = -2

	ROLE_USER  = 1
	ROLE_STAFF = 2
	ROLE_ADMIN = 3
)

type Users struct {
	common.SQLModel
	Name     string `json:"name" gorm:"column:name;"`
	UserName string `json:"user_name" gorm:"column:user_name;"`
	PassWord string `json:"-" gorm:"column:password;"`
	Salt     string `json:"-" gorm:"column:salt"`
	Status   int    `json:"status" gorm:"column:status;"`
	Phone    string `json:"phone" gorm:"column:phone"`
	Role     int    `json:"role" gorm:"column:role"`
	AvatarId int    `json:"-" gorm:"column:avatar_id"`
}

func (Users) TableUsers() string {
	return "users"
}
