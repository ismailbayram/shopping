package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFileStorage_Upload(t *testing.T) {
	fileStorage := NewFileStorage("test_media", "http://localhost:8080/media/")
	defer os.RemoveAll(fileStorage.baseDir)

	_, err := os.Stat(fileStorage.baseDir)
	assert.Nil(t, err)

	fName, err := fileStorage.Upload("test.txt", []byte("Hello World!"))
	assert.Nil(t, err)
	assert.Contains(t, fName, "test")
	assert.Contains(t, fName, ".txt")
}

func TestFileStorage_Url(t *testing.T) {
	fileStorage := NewFileStorage("test_media", "http://localhost:8080/media/")
	defer os.RemoveAll(fileStorage.baseDir)

	_, err := os.Stat(fileStorage.baseDir)
	assert.Nil(t, err)

	fName, err := fileStorage.Upload("test.txt", []byte("Hello World!"))
	assert.Nil(t, err)
	assert.Contains(t, fName, "test")
	assert.Contains(t, fName, ".txt")

	url := fileStorage.Url(fName)
	assert.Equal(t, "http://localhost:8080/media/"+fName, url)
}
