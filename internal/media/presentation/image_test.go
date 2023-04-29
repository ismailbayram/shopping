package presentation

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/ismailbayram/shopping/internal/media/models"
	"github.com/ismailbayram/shopping/test/mocks"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMediaViews_ImageDetailView(t *testing.T) {
	mockIS := &mocks.ImageService{}
	views := NewMediaViews(mockIS)
	var payload map[string]any
	gin.SetMode(gin.TestMode)
	viper.Set("server.domain", "https://shopping.com")
	viper.Set("server.mediaurl", "media")

	// without imageId
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	views.ImageDetailView(ctx)
	assert.Equal(t, http.StatusNotFound, w.Code)
	resp, _ := io.ReadAll(w.Body)
	_ = json.Unmarshal(resp, &payload)
	assert.Equal(t, map[string]any{}, payload)

	// with imageId does not exist
	w = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(w)
	ctx.Params = []gin.Param{{Key: "imageId", Value: "2"}}
	mockIS.On("GetByID", uint(2)).Return(models.Image{}, models.ErrorImageNotFound).Once()
	views.ImageDetailView(ctx)
	assert.Equal(t, http.StatusNotFound, w.Code)
	resp, _ = io.ReadAll(w.Body)
	_ = json.Unmarshal(resp, &payload)
	assert.Equal(t, map[string]any{}, payload)

	// with imageId exists
	w = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(w)
	ctx.Params = []gin.Param{{Key: "imageId", Value: "1"}}
	mockIS.On("GetByID", uint(1)).Return(models.Image{ID: 1, Path: "image.png"}, nil).Once()
	views.ImageDetailView(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
	resp, _ = io.ReadAll(w.Body)
	var imageDTO ImageDTO
	_ = json.Unmarshal(resp, &imageDTO)
	assert.Equal(t, 1, imageDTO.ID)
	assert.Equal(t, "https://shopping.com/media/image.png", imageDTO.Url)
}

func TestMediaViews_ImageCreateView_Fail(t *testing.T) {
	mockIS := &mocks.ImageService{}
	views := NewMediaViews(mockIS)
	var payload map[string]any
	gin.SetMode(gin.TestMode)

	// wrong request type
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("POST", "/api/images/", nil)

	views.ImageCreateView(ctx)
	assert.Equal(t, http.StatusUnsupportedMediaType, w.Code)

	// without file
	w = httptest.NewRecorder()
	buf := new(bytes.Buffer)
	ctx, _ = gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("POST", "/api/images/", nil)
	mw := multipart.NewWriter(buf)
	ctx.Request.Header.Set("Content-Type", mw.FormDataContentType())

	views.ImageCreateView(ctx)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	resp, _ := io.ReadAll(w.Body)
	_ = json.Unmarshal(resp, &payload)
	assert.Equal(t, "'image' file field is required.", payload["image"])

	// with file
	buf = new(bytes.Buffer)
	mw = multipart.NewWriter(buf)
	mw.CreateFormFile("image", "test.png")
	mw.Close()
	w = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("POST", "/api/images/", buf)
	ctx.Request.Header.Set("Content-Type", mw.FormDataContentType())
	views.ImageCreateView(ctx)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestMediaViews_ImageCreateView_Success(t *testing.T) {
	mockIS := &mocks.ImageService{}
	views := NewMediaViews(mockIS)
	gin.SetMode(gin.TestMode)
	viper.Set("server.domain", "https://shopping.com")
	viper.Set("server.mediaurl", "media")

	buf := new(bytes.Buffer)
	mw := multipart.NewWriter(buf)
	fileWriter, _ := mw.CreateFormFile("image", "test.png")
	fileWriter.Write([]byte("image content"))
	mw.Close()
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("POST", "/api/images/", buf)
	ctx.Request.Header.Set("Content-Type", mw.FormDataContentType())

	mockIS.On(
		"Create",
		"test.png",
		[]byte("image content"),
	).Return(
		models.Image{ID: 1, Path: "image.png"},
		nil,
	).Once()
	views.ImageCreateView(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
	resp, _ := io.ReadAll(w.Body)
	var imageDTO ImageDTO
	_ = json.Unmarshal(resp, &imageDTO)
	assert.Equal(t, 1, imageDTO.ID)
	assert.Equal(t, "https://shopping.com/media/image.png", imageDTO.Url)
}
