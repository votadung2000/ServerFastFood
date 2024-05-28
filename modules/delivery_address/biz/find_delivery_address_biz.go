package bizDeliveryAddress

import (
	"context"
	"fastFood/common"
	modelDeliveryAddress "fastFood/modules/delivery_address/model"
)

type FindDeliveryAddressStorage interface {
	FindDeliveryAddress(ctx context.Context, cond map[string]interface{}) (*modelDeliveryAddress.DeliveryAddress, error)
}

type findDeliveryAddressBiz struct {
	store FindDeliveryAddressStorage
}

func NewFindDeliveryAddressBiz(store FindDeliveryAddressStorage) *findDeliveryAddressBiz {
	return &findDeliveryAddressBiz{store: store}
}

func (biz *findDeliveryAddressBiz) FindDeliveryAddress(ctx context.Context, deliveryAddressId, userId int) (*modelDeliveryAddress.DeliveryAddress, error) {
	data, err := biz.store.FindDeliveryAddress(ctx, map[string]interface{}{"id": deliveryAddressId, "user_id": userId})
	if err != nil {
		return nil, common.ErrCannotGetEntity(modelDeliveryAddress.EntityName, err)
	}

	return data, nil
}
