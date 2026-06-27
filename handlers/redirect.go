package handlers

import (
	"database/sql"
	"net/http"

	"shortener/database"

	"github.com/gin-gonic/gin"
)

func Redirect(c *gin.Context) {

	code := c.Param("code")

	var url string

	err := database.DB.QueryRow(
		"SELECT url FROM links WHERE code = ?",
		code,
	).Scan(&url)

	if err == sql.ErrNoRows {
		c.String(http.StatusNotFound, "Link não encontrado")
		return
	}

	if err != nil {
		c.String(http.StatusInternalServerError, "Erro interno")
		return
	}

	c.Redirect(http.StatusFound, url)
}