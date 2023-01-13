package infrastructure

import (
	domain "github.com/ismailbayram/shopping/internal/media/domain/models"
	"gorm.io/gorm"
	"time"
)

type ImageDB struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
	Path      string    `gorm:"not null"`
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

func (idr ImageDBRepository) Create(image domain.Image) (domain.Image, error) {
	imageDB := ImageDB{
		Path: image.Path,
	}

	result := idr.db.Create(&imageDB)
	if result.Error != nil {
		return domain.Image{}, domain.ErrorGeneral
	}

	return domain.Image{
		ID:   imageDB.ID,
		Path: imageDB.Path,
	}, nil
}

func (idr ImageDBRepository) GetByID(id uint) (domain.Image, error) {
	var imageDB ImageDB
	result := idr.db.Where("id = ?", id).First(&imageDB)

	if result.Error != nil {
		return domain.Image{}, domain.ErrorImageNotFound
	}

	return domain.Image{
		ID:   imageDB.ID,
		Path: imageDB.Path,
	}, nil
}
