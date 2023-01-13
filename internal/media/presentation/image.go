package presentation

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ismailbayram/shopping/internal/application"
	domain "github.com/ismailbayram/shopping/internal/media/domain/models"
	"log"
	"net/http"
	"strconv"
)

func ImageDetailView(app *application.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("imageId"))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{})
			return
		}

		image, err := app.Media.Service.GetByID(uint(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"id":  image.ID,
			"url": app.GetMediaUrl(image.Path),
		})
	}
}

func ImageCreateView(app *application.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: check file extension
		// ErrMissingFile
		formFile, err := ctx.FormFile("image")
		if errors.Is(err, http.ErrNotMultipart) {
			ctx.JSON(http.StatusUnsupportedMediaType, gin.H{
				"error": "Content type is not multipart/form-data",
			})
			return
		}
		if errors.Is(err, http.ErrMissingFile) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"image": "'image' file field is required.",
			})
			return
		}

		//formFile, err := ctx.FormFile("image")
		file, err := formFile.Open()
		defer file.Close()
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": domain.ErrorGeneral,
			})
			return
		}

		content := make([]byte, formFile.Size)
		_, err = file.Read(content)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": domain.ErrorGeneral,
			})
			return
		}

		image, err := app.Media.Service.Create(formFile.Filename, content)

		ctx.JSON(http.StatusOK, gin.H{
			"id":  image.ID,
			"url": app.GetMediaUrl(image.Path),
		})
	}
}
