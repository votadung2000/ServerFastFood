package bizUser

import (
	"context"
	"fastFood/common"
	modelUser "fastFood/modules/user/model"
)

type UpdateUserStorage interface {
	FindUser(ctx context.Context, cond map[string]interface{}) (*modelUser.User, error)
	UpdateUser(ctx context.Context, cond map[string]interface{}, dataUpdate *modelUser.UserUpdate) error
}

type updateUserBiz struct {
	store UpdateUserStorage
}

func NewUpdateUserBiz(store UpdateUserStorage) *updateUserBiz {
	return &updateUserBiz{store: store}
}

func (biz *updateUserBiz) UpdateUser(ctx context.Context, id int, dataUpdate *modelUser.UserUpdate) error {
	if err := dataUpdate.Validate(); err != nil {
		return err
	}

	data, err := biz.store.FindUser(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(modelUser.EntityName, err)
	}

	if data.Status != 0 && data.Status == modelUser.STATUS_DELETED {
		return modelUser.ErrUserHasBeenDeleted()
	}

	if data.Status != 0 && data.Status == modelUser.STATUS_BLOCK {
		return modelUser.ErrUserHasBeenBlocked()
	}

	if err := biz.store.UpdateUser(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(modelUser.EntityName, err)
	}

	return nil
}
