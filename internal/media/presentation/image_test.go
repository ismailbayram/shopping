package presentation

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	domain "github.com/ismailbayram/shopping/internal/media/domain/models"
	"github.com/ismailbayram/shopping/test/mocks"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestImageDetailView(t *testing.T) {
	mockIS := &mocks.ImageService{}
	views := NewMediaViews(mockIS)
	var payload map[string]any
	gin.SetMode(gin.TestMode)

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
	mockIS.On("GetByID", uint(2)).Return(domain.Image{}, domain.ErrorImageNotFound).Once()
	views.ImageDetailView(ctx)
	assert.Equal(t, http.StatusNotFound, w.Code)
	resp, _ = io.ReadAll(w.Body)
	_ = json.Unmarshal(resp, &payload)
	assert.Equal(t, map[string]any{}, payload)

	// with imageId exists
	viper.Set("server.domain", "https://shopping.com")
	viper.Set("server.mediaurl", "media")
	w = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(w)
	ctx.Params = []gin.Param{{Key: "imageId", Value: "1"}}
	mockIS.On("GetByID", uint(1)).Return(domain.Image{ID: 1, Path: "image.png"}, nil).Once()
	views.ImageDetailView(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
	resp, _ = io.ReadAll(w.Body)
	_ = json.Unmarshal(resp, &payload)
	assert.Equal(t, 1, int(payload["id"].(float64)))
	assert.Equal(t, "https://shopping.com/media/image.png", payload["url"])
}
