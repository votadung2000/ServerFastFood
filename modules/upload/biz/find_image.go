package bizUpload

import (
	"context"
	"fastFood/common"
	modelUpload "fastFood/modules/upload/model"
)

type FindImageStorage interface {
	FindImage(ctx context.Context, cond map[string]interface{}) (*common.Image, error)
}

type findImageBiz struct {
	store FindImageStorage
}

func NewFindImageBiz(store FindImageStorage) *findImageBiz {
	return &findImageBiz{store: store}
}

func (biz *findImageBiz) FindImage(ctx context.Context, id int) (*common.Image, error) {
	data, err := biz.store.FindImage(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, modelUpload.ErrCannotGetEntity(err)
	}

	return data, nil
}
