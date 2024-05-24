package bizDeliveryAddress

import (
	"context"
	"fastFood/common"
	modelDeliveryAddress "fastFood/modules/delivery_address/model"
)

type ListDeliveryAddressStorage interface {
	ListDeliveryAddress(
		ctx context.Context,
		filter *modelDeliveryAddress.Filter,
		paging *common.Paging,
		cond map[string]interface{},
		moreKeys ...string,
	) ([]modelDeliveryAddress.DeliveryAddress, error)
}

type listDeliveryAddress struct {
	store ListDeliveryAddressStorage
}

func NewListDeliveryAddressBiz(store ListDeliveryAddressStorage) *listDeliveryAddress {
	return &listDeliveryAddress{store: store}
}

func (biz *listDeliveryAddress) ListDeliveryAddress(
	ctx context.Context,
	filter *modelDeliveryAddress.Filter,
	paging *common.Paging,
	id int,
) ([]modelDeliveryAddress.DeliveryAddress, error) {
	data, err := biz.store.ListDeliveryAddress(ctx, filter, paging, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return data, nil
}
