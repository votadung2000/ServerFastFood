package bizDeliveryAddress

import (
	"context"
	"fastFood/common"
	modelDeliveryAddress "fastFood/modules/delivery_address/model"
)

type FindDeliveryAddressDefaultStorage interface {
	FindDeliveryAddress(ctx context.Context, cond map[string]interface{}) (*modelDeliveryAddress.DeliveryAddress, error)
}

type findDeliveryAddressDefaultBiz struct {
	store FindDeliveryAddressDefaultStorage
}

func NewFindDeliveryAddressDefaultBiz(store FindDeliveryAddressDefaultStorage) *findDeliveryAddressDefaultBiz {
	return &findDeliveryAddressDefaultBiz{store: store}
}

func (biz *findDeliveryAddressDefaultBiz) FindDeliveryAddressDefault(ctx context.Context) (*modelDeliveryAddress.DeliveryAddress, error) {
	data, err := biz.store.FindDeliveryAddress(ctx, map[string]interface{}{"default": modelDeliveryAddress.DEFAULT})
	if err != nil {
		return nil, common.ErrCannotGetEntity(modelDeliveryAddress.EntityName, err)
	}

	return data, nil
}
