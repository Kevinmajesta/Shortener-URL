package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/Kevinmajesta/Shortener-URL/database"
	"github.com/Kevinmajesta/Shortener-URL/models"
	"github.com/gin-gonic/gin"
)

// Fungsi untuk generate short URL unik
func GenerateShortURL() string {
	b := make([]byte, 4)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:6]
}

func CreateShortURL(c *gin.Context) {
	var requestBody struct {
		LongURL string `json:"long_url"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	shortURL := GenerateShortURL()
	url := models.URL{ShortURL: shortURL, LongURL: requestBody.LongURL}

	database.DB.Create(&url) // ✅ Gunakan database.DB

	c.JSON(http.StatusOK, gin.H{"short_url": "http://localhost:8080/" + shortURL})
}

// Endpoint untuk redirect dari short URL ke long URL
func RedirectShortURL(c *gin.Context) {
	shortURL := c.Param("shortURL")
	var url models.URL

	result := database.DB.Where("short_url = ?", shortURL).First(&url) // ✅ Gunakan database.DB
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusFound, url.LongURL)
}
