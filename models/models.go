package models

type FormatGetList struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

type FormatGetListProducts struct {
	FormatGetList
	CategoryId int    `json:"category_id" form:"category_id"`
	Name       string `json:"name" form:"name"`
}

type FormatResponse struct {
	Total int64       `json:"total" form:"total"`
	Data  interface{} `json:"data" form:"data"`
}

type FormatGetFavorites struct {
	FormatGetList
	UserId    int `json:"user_id" gorm:"column:user_id;"`
	ProductId int `json:"product_id" gorm:"column:product_id;"`
}
