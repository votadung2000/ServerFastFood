package modelDeliveryAddress

import (
	"errors"
	"fastFood/common"
)

func ErrValidateRequest(msg, key string) *common.AppError {
	return common.ErrorResponse(errors.New(msg), msg, msg, key)
}

func ErrDeliveryAddressHasBeenDeleted() *common.AppError {
	return common.CustomError(
		errors.New("delivery address has been deleted"),
		"delivery address has been deleted",
		"ERR_DELIVERY_ADDRESS_HAS_BEEN_DELETED",
	)
}
