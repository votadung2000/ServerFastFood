package modelCategory

import (
	"errors"
	"fastFood/common"
)

func ErrValidateRequest(msg, key string) *common.AppError {
	return common.ErrorResponse(errors.New(msg), msg, msg, key)
}

func ErrCategoryHasBeenDeleted() *common.AppError {
	return common.CustomError(
		errors.New("the category has been deleted"),
		"the category has been deleted",
		"ERR_CATEGORY_HAS_BEEN_DELETED",
	)
}

func ErrCategoryHasBeenBlocked() *common.AppError {
	return common.CustomError(
		errors.New("the category has been blocked"),
		"the category has been blocked",
		"ERR_CATEGORY_HAS_BEEN_BLOCKED",
	)
}
