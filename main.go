package main

import (
	"github.com/Kevinmajesta/Shortener-URL/database"
	"github.com/Kevinmajesta/Shortener-URL/handlers"
	"github.com/Kevinmajesta/Shortener-URL/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// ✅ Panggil fungsi ConnectDatabase sebelum menggunakan database
	database.ConnectDatabase()

	// ✅ Pastikan migrasi database berjalan
	database.DB.AutoMigrate(&models.URL{})

	r := gin.Default()

	// Routes
	r.POST("/shorten", handlers.CreateShortURL)
	r.GET("/:shortURL", handlers.RedirectShortURL)

	// Jalankan server di port 8080
	r.Run(":8080")
}
