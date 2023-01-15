package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFileStorage(t *testing.T) {
	fileStorage := NewFileStorage("test_media")
	defer os.RemoveAll(fileStorage.baseDir)

	_, err := os.Stat(fileStorage.baseDir)
	assert.Nil(t, err)

	fName, err := fileStorage.Upload("test.txt", []byte("Hello World!"))
	assert.Nil(t, err)
	assert.Contains(t, fName, "test")
	assert.Contains(t, fName, ".txt")
}
