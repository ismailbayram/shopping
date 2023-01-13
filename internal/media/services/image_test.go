package services

import (
	"errors"
	domain "github.com/ismailbayram/shopping/internal/media/domain/models"
	"github.com/ismailbayram/shopping/test/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImageService_GetByID(t *testing.T) {
	image := domain.Image{
		ID:   2,
		Path: "media/images/test.png",
	}

	mockedIR := &mocks.ImageRepository{}
	mockedIS := &mocks.ImageStorage{}
	IS := NewImageService(mockedIR, mockedIS)

	mockedIR.On("GetByID", uint(1)).Return(domain.Image{}, domain.ErrorImageNotFound)
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
	assert.Equal(t, domain.ErrorGeneral, err)

	mockedIS.On("Upload", "image.png", []byte("file content")).Return("images/image.png", nil).Once()
	mockedIR.On("Create", domain.Image{Path: "images/image.png"}).Return(domain.Image{ID: 1, Path: "images/image.png"}, nil)
	image, err = IS.Create("image.png", []byte("file content"))
	assert.Nil(t, err)
	assert.Equal(t, uint(1), image.ID)
	assert.Equal(t, "images/image.png", image.Path)
}
