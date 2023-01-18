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

	api := r.Group("/api")
	api.GET("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	api.GET("/images/:imageId", app.Media.Views.ImageDetailView)
	api.POST("/images", app.Media.Views.ImageCreateView)

	return r
}
