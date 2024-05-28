package bizDeliveryAddress

import (
	"context"
	"fastFood/common"
	modelDeliveryAddress "fastFood/modules/delivery_address/model"
)

type DeleteDeliveryAddressStorage interface {
	FindDeliveryAddress(ctx context.Context, cond map[string]interface{}) (*modelDeliveryAddress.DeliveryAddress, error)
	DeleteDeliveryAddress(ctx context.Context, cond map[string]interface{}) error
}

type deleteDeliveryAddressBiz struct {
	store DeleteDeliveryAddressStorage
}

func DeleteDeliveryAddressBiz(store DeleteDeliveryAddressStorage) *deleteDeliveryAddressBiz {
	return &deleteDeliveryAddressBiz{store: store}
}

func (biz *deleteDeliveryAddressBiz) DeleteDeliveryAddress(ctx context.Context, id int) error {
	address, err := biz.store.FindDeliveryAddress(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(modelDeliveryAddress.EntityName, err)
	}

	if address == nil {
		return modelDeliveryAddress.ErrDeliveryAddressHasBeenDeleted()
	}

	if err := biz.store.DeleteDeliveryAddress(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}

	return nil
}
