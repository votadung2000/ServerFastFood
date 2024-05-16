package bizOrder

import (
	"context"
	"fastFood/common"
	modelOrder "fastFood/modules/order/model"
)

type ListOrderStorage interface {
	ListOrder(
		ctx context.Context,
		filter *modelOrder.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]modelOrder.Order, error)
}

type listOrderBiz struct {
	store ListOrderStorage
}

func NewListOrderBiz(store ListOrderStorage) *listOrderBiz {
	return &listOrderBiz{store: store}
}

func (biz *listOrderBiz) ListOrder(
	ctx context.Context,
	filter *modelOrder.Filter,
	paging *common.Paging,
) ([]modelOrder.Order, error) {
	data, err := biz.store.ListOrder(ctx, filter, paging)
	if err != nil {
		return nil, err
	}

	return data, nil
}
