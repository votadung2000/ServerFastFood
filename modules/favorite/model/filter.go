package modelFavorite

type Filter struct {
	UserId int `json:"user_id" gorm:"column:user_id;"`
}
