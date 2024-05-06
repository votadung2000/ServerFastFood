package modelUser

import (
	"errors"
	"fastFood/common"
)

func ErrValidateRequest(msg, key string) *common.AppError {
	return common.ErrorResponse(errors.New(msg), msg, msg, key)
}

func ErrUserNameExisted() *common.AppError {
	return common.CustomError(
		errors.New("user name has already existed"),
		"User name has already existed",
		"ERR_EMAIL_EXISTED",
	)
}

func ErrCannotCreateEntity(err error) *common.AppError {
	return common.CustomError(
		err,
		"Cannot create user",
		"ERR_CANNOT_CREATE_USER",
	)
}

func ErrCannotChangePassword(err error) *common.AppError {
	return common.CustomError(
		err,
		"Cannot change password",
		"ERR_CANNOT_CHANGE_PASSWORD",
	)
}

func ErrUserNameOrPasswordInvalid() *common.AppError {
	return common.CustomError(
		errors.New("user name or password invalid"),
		"User name or password invalid",
		"ERR_USER_NAME_OR_PASS_INVALID",
	)
}

func ErrUserHasBeenDeleted() *common.AppError {
	return common.CustomError(
		errors.New("the user has been deleted"),
		"the user has been deleted",
		"ERR_USER_HAS_BEEN_DELETED",
	)
}

func ErrUserHasBeenBlocked() *common.AppError {
	return common.CustomError(
		errors.New("the user has been blocked"),
		"the user has been blocked",
		"ERR_USER_HAS_BEEN_BLOCKED",
	)
}
