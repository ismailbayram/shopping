package services

import (
	"errors"
	"github.com/ismailbayram/shopping/internal/media/models"
	"github.com/ismailbayram/shopping/test/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImageService_GetByID(t *testing.T) {
	image := models.Image{
		ID:   2,
		Path: "media/images/test.png",
	}

	mockedIR := &mocks.ImageRepository{}
	mockedIS := &mocks.ImageStorage{}
	IS := NewImageService(mockedIR, mockedIS)

	mockedIR.On("GetByID", uint(1)).Return(models.Image{}, models.ErrorImageNotFound)
	mockedIS.On("Url", "media/images/test.png").Return("http://localhost/media/images/test.png")
	imageGot, err := IS.GetByID(1)
	assert.Equal(t, uint(0), imageGot.ID)
	assert.NotNil(t, err)

	mockedIR.On("GetByID", uint(2)).Return(image, nil)
	imageGot, err = IS.GetByID(2)
	assert.NotEqual(t, uint(0), imageGot)
	assert.Nil(t, err)
	assert.Equal(t, image.ID, imageGot.ID)
}

func TestImageService_Create(t *testing.T) {
	mockedIR := &mocks.ImageRepository{}
	mockedIS := &mocks.ImageStorage{}
	IS := NewImageService(mockedIR, mockedIS)

	mockedIS.On("Upload", "image.png", []byte("file content")).Return("", errors.New("dump")).Once()
	image, err := IS.Create("image.png", []byte("file content"))
	assert.Equal(t, uint(0), image.ID)
	assert.Equal(t, models.ErrorGeneral, err)

	mockedIS.On("Upload", "image.png", []byte("file content")).Return("images/image.png", nil).Once()
	mockedIR.On("Create", models.Image{Path: "images/image.png"}).Return(models.Image{ID: 1, Path: "images/image.png"}, nil)
	mockedIS.On("Url", "images/image.png").Return("http://localhost/media/images/test.png")
	image, err = IS.Create("image.png", []byte("file content"))
	assert.Nil(t, err)
	assert.Equal(t, uint(1), image.ID)
	assert.Equal(t, "images/image.png", image.Path)
}
