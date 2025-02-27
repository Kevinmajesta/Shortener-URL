package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Variabel global untuk menyimpan koneksi database

func ConnectDatabase() {
	dsn := "host=- user=- password=- dbname=postgres port=5432 sslmode=require"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Gagal terhubung ke database:", err)
	}

	DB = database
	log.Println("✅ Database connected!")
}
