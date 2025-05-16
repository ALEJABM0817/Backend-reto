package database

import (
	"fmt"
	"log"
	"os"

	"github.com/ALEJABM0817/TGolang/models" // <-- Importa tu paquete de modelos
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Fail to load file .env")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error to conect database: ", err)
	}

	DB = db
	log.Println("Conexion successful")

	if err := DB.AutoMigrate(&models.AnalystRating{}); err != nil {
		log.Fatal("Table migration error: ", err)
	}
}
