package tokenProvider

import (
	"errors"
	"fastFood/common"
)

func ErrNotFound() *common.AppError {
	return common.CustomError(
		errors.New("Token not found"),
		"Token not found",
		"ERR_TOKEN_NOT_FOUND",
	)
}

func ErrUnexpectedSigningMethod(err string) *common.AppError {
	return common.CustomError(
		errors.New(err),
		"unexpected signing method",
		"ERR_UNEXPECTED_SIGNING_METHOD",
	)
}
