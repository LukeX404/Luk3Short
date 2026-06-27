package main

import (
	"shortener/database"
	"shortener/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Init()

	r := gin.Default()

	r.LoadHTMLFiles("index.html")

	routes.Register(r)

	r.Run(":8080")
}