package bizDeliveryAddress

import (
	"context"
	modelDeliveryAddress "fastFood/modules/delivery_address/model"
)

type CreateDeliveryAddressStorage interface {
	CreateDeliveryAddress(ctx context.Context, data *modelDeliveryAddress.CreateDeliveryAddress) error
}

type createDeliveryAddressBiz struct {
	store CreateDeliveryAddressStorage
}

func NewCreateDeliveryAddressBiz(store CreateDeliveryAddressStorage) *createDeliveryAddressBiz {
	return &createDeliveryAddressBiz{store: store}
}

func (biz *createDeliveryAddressBiz) CreateDeliveryAddress(ctx context.Context, data *modelDeliveryAddress.CreateDeliveryAddress) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.CreateDeliveryAddress(ctx, data); err != nil {
		return err
	}

	return nil
}
