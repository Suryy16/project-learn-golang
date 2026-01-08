package models

import (
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Drop tables jika sudah ada (HATI-HATI: Data akan hilang!)
	//db.Migrator().DropTable(&Todo{}, &User{})

	// AutoMigrate akan membuat tabel users dan todos dengan kolom yang benar
	err = db.AutoMigrate(&User{}, &Todo{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Successfully Connected to Database")
	DB = db
}
