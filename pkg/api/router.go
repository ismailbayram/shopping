package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ismailbayram/shopping/internal/users"
	"net/http"
)

type App struct {
	Users users.Users
}

func NewRouter(app *App) *gin.Engine {
	//f, _ := os.Create("shopping.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//	r.Use(gin.Recovery())

	r := gin.Default()
	api := r.Group("/api")

	api.GET("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r
}
