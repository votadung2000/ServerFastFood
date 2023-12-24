package modelProduct

type Filter struct {
	Name       string `json:"name,omitempty" form:"name"`
	Status     int    `json:"status,omitempty" form:"status"`
	CategoryId int    `json:"category_id,omitempty" form:"category_id"`
}
