package bizUser

import (
	"context"
	"fastFood/common"
	modelUser "fastFood/modules/user/model"
)

type UpdatePasswordStorage interface {
	FindUser(ctx context.Context, cond map[string]interface{}) (*modelUser.User, error)
	UpdatePassword(ctx context.Context, cond map[string]interface{}, dataUpdate *modelUser.UpdatePassword) error
}

type updatePasswordBiz struct {
	store   UpdatePasswordStorage
	hashery Hashery
	expiry  int
}

func NewUpdatePasswordBiz(store UpdatePasswordStorage, hashery Hashery, expiry int) *updatePasswordBiz {
	return &updatePasswordBiz{store: store, hashery: hashery, expiry: expiry}
}

func (biz *updatePasswordBiz) UpdatePassword(ctx context.Context, data *modelUser.UpdatePassword) error {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return modelUser.ErrUserNameOrPasswordInvalid()
	}

	passHashed := biz.hashery.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return modelUser.ErrUserNameOrPasswordInvalid()
	}

	salt := common.GenSalt(50)
	data.Salt = salt
	data.Password = biz.hashery.Hash(data.NewPassword + salt)
	data.Email = ""
	data.NewPassword = ""

	if err := biz.store.UpdatePassword(ctx, map[string]interface{}{"id": user.Id}, data); err != nil {
		return modelUser.ErrCannotChangePassword(err)
	}

	return nil
}
