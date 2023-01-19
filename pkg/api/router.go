package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ismailbayram/shopping/internal/application"
	"net/http"
)

func NewRouter(app *application.Application) *gin.Engine {
	//f, _ := os.Create("shopping.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//	r.Use(gin.Recovery())

	r := gin.Default()

	r.StaticFS(app.MediaUrl, http.Dir("media"))

	adminAPI := r.Group("/admin/api")
	adminAPI.GET("/images/:imageId", app.Media.Views.ImageDetailView)
	adminAPI.POST("/images", app.Media.Views.ImageCreateView)

	return r
}
