package bizUpload

import (
	"bytes"
	"context"
	"fastFood/common"
	modelUpload "fastFood/modules/upload/model"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"path/filepath"
)

type InsertImageStorage interface {
	InsertImage(ctx context.Context, data *common.Image) error
}

type createImageBiz struct {
	store InsertImageStorage
}

func NewCreateImageBiz(store InsertImageStorage) *createImageBiz {
	return &createImageBiz{store: store}
}

func (biz *createImageBiz) CreateImage(ctx context.Context, data []byte, fileName, dstLocal string, img *common.Image) error {
	fileBytes := bytes.NewReader(data)

	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		errDelete := deleteFile(dstLocal)

		if errDelete != nil {
			return modelUpload.ErrCannotSaveFile(err)
		}

		return modelUpload.ErrFileIsNotImage(err)
	}

	fileExt := filepath.Ext(fileName)

	img.Url = dstLocal
	img.Width = w
	img.Height = h
	img.CloudName = "local"
	img.Extension = fileExt

	domain := os.Getenv("DOMAIN")
	img.Fulfil(domain)

	if err := biz.store.InsertImage(ctx, img); err != nil {
		return modelUpload.ErrCannotCreateEntity(err)
	}

	return nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}

func deleteFile(path string) error {
	err := os.Remove(path)

	if err != nil {
		return err
	}

	return nil
}
