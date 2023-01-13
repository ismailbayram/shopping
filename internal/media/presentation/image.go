package presentation

import (
	"github.com/gin-gonic/gin"
	"github.com/ismailbayram/shopping/internal/application"
	"net/http"
)

func ImageDetailView(app *application.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"test": "Hello World!",
		})
	}
}
