package handlers

import (
	"net/http"
	"strings"

	"shortener/database"
	"shortener/utils"

	"github.com/gin-gonic/gin"
)

func Shorten(c *gin.Context) {

	var body struct {
		URL string `json:"url"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "URL inválida",
		})
		return
	}

	body.URL = strings.TrimSpace(body.URL)

	if body.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Informe uma URL",
		})
		return
	}

	if !strings.HasPrefix(body.URL, "http://") &&
		!strings.HasPrefix(body.URL, "https://") {
		body.URL = "https://" + body.URL
	}

	// Verifica se a URL já existe
	var existingCode string

	err := database.DB.QueryRow(
		"SELECT code FROM links WHERE url = ?",
		body.URL,
	).Scan(&existingCode)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"short": "http://localhost:8080/" + existingCode,
		})
		return
	}

	// Gera um código único
	var code string

	for {

		code = utils.RandomString()

		var exists int

		database.DB.QueryRow(
			"SELECT COUNT(*) FROM links WHERE code = ?",
			code,
		).Scan(&exists)

		if exists == 0 {
			break
		}
	}

	_, err = database.DB.Exec(
		"INSERT INTO links(code, url) VALUES(?, ?)",
		code,
		body.URL,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"short": "http://localhost:8080/" + code,
	})
}