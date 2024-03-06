package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

const (
	TYPE_IMG_PROFILE  = 1
	TYPE_IMG_CATEGORY = 2
	TYPE_IMG_PRODUCT  = 3
	TYPE_IMG_OTHER    = 0
)

type Image struct {
	SQLModel
	Url       string `json:"url" gorm:"column:url"`
	Width     int    `json:"width" gorm:"column:width"`
	Height    int    `json:"height" gorm:"column:height"`
	Type      int    `json:"type" gorm:"column:type"`
	CloudName string `json:"cloud_name,omitempty" gorm:"column:cloud_name"`
	Extension string `json:"extension,omitempty" gorm:"column:extension"`
}

func (Image) TableName() string {
	return "images"
}

func (i *Image) Fulfil(domain string) {
	i.Url = fmt.Sprintf("%s/%s", domain, i.Url)
}

// get interface of DB ->
func (i *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*i = img

	return nil
}

// struct -> DB
func (i *Image) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}

	return json.Marshal(i)
}
