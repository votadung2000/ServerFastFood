package bizDeliveryAddress

import (
	"context"
	"fastFood/common"
	modelDeliveryAddress "fastFood/modules/delivery_address/model"
)

type UpdateDeliveryAddressStorage interface {
	FindDeliveryAddress(ctx context.Context, cond map[string]interface{}) (*modelDeliveryAddress.DeliveryAddress, error)
	UpdateCategory(
		ctx context.Context,
		cond map[string]interface{},
		dataUpdate *modelDeliveryAddress.DeliveryAddressUpdate,
	) error
}

type updateDeliveryAddressBiz struct {
	store UpdateDeliveryAddressStorage
}

func NewUpdateDeliveryAddressBiz(store UpdateDeliveryAddressStorage) *updateDeliveryAddressBiz {
	return &updateDeliveryAddressBiz{store: store}
}

func (biz *updateDeliveryAddressBiz) UpdateDeliveryAddress(
	ctx context.Context,
	id int,
	data *modelDeliveryAddress.DeliveryAddressUpdate,
) error {
	address, err := biz.store.FindDeliveryAddress(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(modelDeliveryAddress.EntityName, err)
	}

	if address == nil {
		return modelDeliveryAddress.ErrDeliveryAddressHasBeenDeleted()
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.UpdateCategory(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return err
	}

	return nil
}
