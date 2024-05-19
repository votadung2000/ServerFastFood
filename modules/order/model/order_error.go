package modelOrder

import (
	"errors"
	"fastFood/common"
)

func ErrValidateRequest(msg, key string) *common.AppError {
	return common.ErrorResponse(errors.New(msg), msg, msg, key)
}

func ErrOrderHasBeenCanceled() *common.AppError {
	return common.CustomError(
		errors.New("the order has been canceled"),
		"the order has been canceled",
		"ERR_ORDER_HAS_BEEN_CANCELED",
	)
}
