package api

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/ismailbayram/shopping/config"
	"github.com/ismailbayram/shopping/internal/application"
)

func NewRouter(app *application.Application, cfg config.ServerConfiguration) *gin.Engine {
	router := gin.Default()

	router.Use(static.Serve(cfg.MediaUrl, static.LocalFile("./media", true)))

	router.Use(gin.Logger())
	router.Use(PanicLoggerMiddleware)
	router.Use(SecurityMiddleware)
	router.Use(ErrorHandlerMiddleware)
	router.Use(AuthenticationMiddleware(app.Users.Service))

	adminAPI := router.Group("/admin/api")
	adminAPI.GET("/images/:imageId", app.Media.Views.ImageDetailView)
	adminAPI.POST("/images", app.Media.Views.ImageCreateView)

	return router
}
