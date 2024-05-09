package modelFavorite

type Filter struct {
	UserId     int `json:"user_id,omitempty" gorm:"column:user_id;"`
	StatusPr   int `json:"status,omitempty" form:"status"`
	CategoryId int `json:"category_id,omitempty" form:"category_id"`
}
