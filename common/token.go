package common

type TokenPayLoad struct {
	UId int `json:"user_id"`
}

func (p TokenPayLoad) UserId() int {
	return p.UId
}
