package modelFavorite

import (
	"errors"
	"time"
)

const (
	STATUS_ACTION  = 1
	STATUS_DELETED = -2
)

var (
	ErrUserIsBlank    = errors.New("the user cannot be blank")
	ErrProductIsBlank = errors.New("the product cannot be blank")
	ErrDeleted        = errors.New("favorite product has been deleted")
)

type Favorite struct {
	Id        int        `json:"id" gorm:"column:id;"`
	UserId    int        `json:"user_id" gorm:"column:user_id;"`
	ProductId int        `json:"product_id" gorm:"column:product_id;"`
	Status    int        `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (Favorite) TableName() string {
	return "favorites"
}

type FavoriteCreate struct {
	UserId    int `json:"user_id" gorm:"column:user_id;"`
	ProductId int `json:"product_id" gorm:"column:product_id;"`
}

func (FavoriteCreate) TableName() string {
	return Favorite{}.TableName()
}
