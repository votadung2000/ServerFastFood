package bizFAQ

import (
	"context"
	"fastFood/common"
	modelFAQ "fastFood/modules/helps_and_faqs/model"
)

type ListFAQStorage interface {
	ListFAQ(
		ctx context.Context,
		filter *modelFAQ.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]modelFAQ.FAQ, error)
}

type listFAQList struct {
	store ListFAQStorage
}

func NewListFAQBiz(store ListFAQStorage) *listFAQList {
	return &listFAQList{store: store}
}

func (biz *listFAQList) ListFAQ(
	ctx context.Context,
	filter *modelFAQ.Filter,
	paging *common.Paging,
) ([]modelFAQ.FAQ, error) {
	data, err := biz.store.ListFAQ(
		ctx,
		filter,
		paging,
	)

	if err != nil {
		return nil, common.ErrCannotListEntity(modelFAQ.EntityName, err)
	}

	return data, nil
}
