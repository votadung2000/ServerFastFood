package bizUser

import (
	"context"
	"fastFood/common"
	modelUser "fastFood/modules/user/model"
)

type InsertUserStorage interface {
	FindUser(ctx context.Context, cond map[string]interface{}) (*modelUser.User, error)
	InsertUser(ctx context.Context, data *modelUser.UserCreate) error
}

type Hashery interface {
	Hash(data string) string
}

type createUserBiz struct {
	store   InsertUserStorage
	hashery Hashery
}

func NewCreateUserBiz(store InsertUserStorage, hashery Hashery) *createUserBiz {
	return &createUserBiz{store: store, hashery: hashery}
}

func (biz *createUserBiz) CreateUser(ctx context.Context, data *modelUser.UserCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	user, _ := biz.store.FindUser(ctx, map[string]interface{}{"user_name": data.UserName})

	if user != nil {
		return modelUser.ErrUserNameExisted()
	}

	salt := common.GenSalt(50)

	data.Salt = salt
	data.Password = biz.hashery.Hash(data.Password + salt)
	data.Role = modelUser.ROLE_USER

	if err := biz.store.InsertUser(ctx, data); err != nil {
		return modelUser.ErrCannotCreateEntity(err)
	}

	return nil
}
