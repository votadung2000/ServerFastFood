package modelProduct

import (
	"errors"
	"fastFood/common"
)

func ErrValidateRequest(msg, key string) *common.AppError {
	return common.ErrorResponse(errors.New(msg), msg, msg, key)
}

func ErrProductHasBeenDeleted() *common.AppError {
	return common.CustomError(
		errors.New("the product has been deleted"),
		"the product has been deleted",
		"ERR_PRODUCT_HAS_BEEN_DELETED",
	)
}
