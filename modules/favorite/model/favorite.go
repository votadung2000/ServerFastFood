package modelFavorite

import (
	"errors"
	"fastFood/common"
	modelProduct "fastFood/modules/product/model"
)

const (
	STATUS_ACTION  = 1
	STATUS_DELETED = -2
)

const (
	EntityName = "Favorite"
)

var (
	ErrUserIsBlank      = errors.New("the user cannot be blank")
	ErrProductIsBlank   = errors.New("the product cannot be blank")
	ErrDeleted          = errors.New("favorite product has been deleted")
	ErrUserIdIsBlank    = "the user id cannot be blank"
	ErrProductIdIsBlank = "the product id cannot be blank"
)

type Favorite struct {
	common.SQLModel
	UserId    int                          `json:"user_id" gorm:"column:user_id;"`
	Status    int                          `json:"status" gorm:"column:status;"`
	ProductId int                          `json:"-" gorm:"column:product_id;"`
	Product   *modelProduct.PreloadProduct `json:"product" gorm:"foreignKey:ProductId"`
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

func (f *FavoriteCreate) Validate() error {
	if f.UserId == 0 {
		return ErrValidateRequest(ErrUserIdIsBlank, "ERR_USER_ID_IS_BLANK")
	}

	if f.ProductId == 0 {
		return ErrValidateRequest(ErrProductIdIsBlank, "ERR_USER_ID_IS_BLANK")
	}

	return nil
}
