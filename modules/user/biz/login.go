package bizUser

import (
	"context"
	"fastFood/common"
	"fastFood/components/tokenProvider"
	modelUser "fastFood/modules/user/model"
)

type FindUserStorage interface {
	FindUser(ctx context.Context, cond map[string]interface{}) (*modelUser.User, error)
}

type loginBiz struct {
	store         FindUserStorage
	hashery       Hashery
	tokenProvider tokenProvider.Provider
	expiry        int
}

func NewLoginBiz(
	store FindUserStorage,
	hashery Hashery,
	tokenProvider tokenProvider.Provider,
	expiry int,
) *loginBiz {
	return &loginBiz{
		store:         store,
		hashery:       hashery,
		tokenProvider: tokenProvider,
		expiry:        expiry,
	}
}

func (biz *loginBiz) Login(ctx context.Context, data *modelUser.Login) (tokenProvider.Token, error) {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"user_name": data.UserName})

	if err != nil {
		return nil, modelUser.ErrUserNameOrPasswordInvalid()
	}

	passHashed := biz.hashery.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, modelUser.ErrUserNameOrPasswordInvalid()
	}

	payload := &common.TokenPayLoad{
		UId: user.Id,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}
