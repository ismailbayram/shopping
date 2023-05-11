package presentation

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (view *MediaViews) ImageDetailView(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("imageId"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}

	image, err := view.Service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, ToImageDTO(image))
}

func (view *MediaViews) ImageCreateView(ctx *gin.Context) {
	// TODO: check file extension, file size
	// TODO: authorization
	formFile, err := ctx.FormFile("image")
	if errors.Is(err, http.ErrNotMultipart) {
		ctx.JSON(http.StatusUnsupportedMediaType, gin.H{
			"error": "Content type is not multipart/form-data",
		})
		return
	}
	if errors.Is(err, http.ErrMissingFile) || fmt.Sprint(err) == "missing form body" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"image": "'image' file field is required.",
		})
		return
	}

	file, err := formFile.Open()
	defer file.Close()
	if err != nil {
		_ = ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	content := make([]byte, formFile.Size)
	_, err = file.Read(content)
	if err != nil {
		_ = ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	image, err := view.Service.Create(formFile.Filename, content)
	if err != nil {
		_ = ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	ctx.JSON(http.StatusOK, ToImageDTO(image))
}
