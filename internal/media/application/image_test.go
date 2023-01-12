package application

import (
	domain "github.com/ismailbayram/shopping/internal/media/domain/models"
	"github.com/ismailbayram/shopping/test/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImageService_GetByID(t *testing.T) {
	image := &domain.Image{
		ID:   2,
		Path: "media/images/test.png",
	}

	mockedIR := &mocks.ImageRepository{}
	PS := NewImageService(mockedIR)

	mockedIR.On("GetByID", uint(1)).Return(nil, domain.ErrorImageNotFound)
	imageGot, err := PS.GetByID(1)
	assert.Nil(t, imageGot)
	assert.NotNil(t, err)

	mockedIR.On("GetByID", uint(2)).Return(image, nil)
	imageGot, err = PS.GetByID(2)
	assert.NotNil(t, imageGot)
	assert.Nil(t, err)
	assert.Equal(t, image.ID, imageGot.ID)
}
