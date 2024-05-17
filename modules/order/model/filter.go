package modelOrder

type Filter struct {
	Status     int  `json:"status,omitempty" form:"status"`
	IsUpcoming bool `json:"is_upcoming,omitempty" form:"is_upcoming"`
	IsHistory  bool `json:"is_history,omitempty" form:"is_history"`
}
