package routes

import (
	"shortener/handlers"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {

	r.GET("/", handlers.Home)

	r.POST("/api/shorten", handlers.Shorten)

	r.GET("/:code", handlers.Redirect)
}