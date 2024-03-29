package infrastructure

import (
	"github.com/ismailbayram/shopping/internal/media/models"
	"github.com/ismailbayram/shopping/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ImageDBTestSuite struct {
	test.AppTestSuite
}

func (s *ImageDBTestSuite) TestCreate() {
	idbr := NewImageDBRepository(s.DB)

	imageDB, err := idbr.Create(models.Image{Path: "test.png"})
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), models.Image{ID: imageDB.ID, Path: "test.png"}, imageDB)

	imageDB, err = idbr.Create(models.Image{Path: "test.png"})
	assert.Equal(s.T(), err, models.ErrorGeneral)
	assert.Equal(s.T(), models.Image{ID: 0, Path: ""}, imageDB)
}

func (s *ImageDBTestSuite) TestGetByID() {
	idbr := NewImageDBRepository(s.DB)

	created, err := idbr.Create(models.Image{Path: "test.png"})
	assert.Nil(s.T(), err)

	imageDB, err := idbr.GetByID(created.ID)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), models.Image{ID: imageDB.ID, Path: "test.png"}, imageDB)

	notExist, err := idbr.GetByID(0)
	assert.Equal(s.T(), models.ErrorImageNotFound, err)
	assert.Equal(s.T(), models.Image{ID: 0, Path: ""}, notExist)
}

func TestImageDBTestSuite(t *testing.T) {
	imageDbTestSuite := new(ImageDBTestSuite)
	imageDbTestSuite.Models = []interface{}{&ImageDB{}}
	suite.Run(t, imageDbTestSuite)
}
