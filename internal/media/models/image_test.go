package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImage_Url(t *testing.T) {
	image := Image{ID: 1, Path: "images/image.png"}
	assert.Equal(t, "http://localhost:8080/images/image.png", image.Url("http://localhost:8080"))
}
