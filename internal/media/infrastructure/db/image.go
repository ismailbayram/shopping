package infrastructure

import (
	"errors"
	"github.com/ismailbayram/shopping/internal/media/models"
	"gorm.io/gorm"
	"log"
	"time"
)

type ImageDB struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
	Path      string    `gorm:"unique;not null;default:null"`
}

func (ImageDB) TableName() string {
	return "media_images"
}

type ImageDBRepository struct {
	db *gorm.DB
}

func NewImageDBRepository(db *gorm.DB) ImageDBRepository {
	return ImageDBRepository{
		db: db,
	}
}

func (idr ImageDBRepository) Create(image models.Image) (models.Image, error) {
	imageDB := ImageDB{
		Path: image.Path,
	}

	result := idr.db.Create(&imageDB)
	if result.Error != nil {
		log.Println(result.Error)
		return models.Image{}, models.ErrorGeneral
	}

	return models.Image{
		ID:   imageDB.ID,
		Path: imageDB.Path,
	}, nil
}

func (idr ImageDBRepository) GetByID(id uint) (models.Image, error) {
	var imageDB ImageDB
	result := idr.db.Where("id = ?", id).First(&imageDB)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Image{}, models.ErrorImageNotFound
		}
		log.Println(result.Error)
		return models.Image{}, models.ErrorGeneral
	}

	return models.Image{
		ID:   imageDB.ID,
		Path: imageDB.Path,
	}, nil
}
