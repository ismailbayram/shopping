package models

import (
	"errors"
	"fmt"
)

var (
	ErrorGeneral       = errors.New("Something went wrong, please try again.")
	ErrorImageNotFound = errors.New("Image not found.")
)

type Image struct {
	ID   uint
	Path string
}

func (i *Image) Url(baseUrl string) string {
	return fmt.Sprintf("%s/%s", baseUrl, i.Path)
}
