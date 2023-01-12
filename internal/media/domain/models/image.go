package domain

import "errors"

var (
	ErrorGeneral       = errors.New("Something went wrong, please try again.")
	ErrorImageNotFound = errors.New("Image not found.")
)

type Image struct {
	ID   uint
	Path string
}
