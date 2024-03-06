package modelUpload

import "fastFood/common"

func ErrCannotCreateEntity(err error) *common.AppError {
	return common.CustomError(
		err,
		"Cannot create image",
		"ERR_CANNOT_CREATE_IMAGE",
	)
}

func ErrCannotGetEntity(err error) *common.AppError {
	return common.CustomError(
		err,
		"Cannot get image",
		"ERR_CANNOT_GET_IMAGE",
	)
}

func ErrFileIsNotImage(err error) *common.AppError {
	return common.CustomError(
		err,
		"File is not image",
		"ERR_FILE_IS_NOT_IMAGE",
	)
}

func ErrCannotSaveFile(err error) *common.AppError {
	return common.CustomError(
		err,
		"Cannot save uploaded file",
		"ERR_CANNOT_SAVE_FILE",
	)
}

func ErrMissingField(err error) *common.AppError {
	return common.CustomError(
		err,
		"missing field parameter",
		"ERR_MISSING_FIELD",
	)
}
