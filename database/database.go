package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ALEJABM0817/TGolang/models" // <-- Importa tu paquete de modelos
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	_ = godotenv.Load()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var db *gorm.DB
	var err error
	maxAttempts := 10

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Intento %d: Error al conectar a la base de datos: %v", attempts, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos después de varios intentos: ", err)
	}

	DB = db
	log.Println("Conexion successful")
	log.Println("DSN:", dsn)

	if err := DB.AutoMigrate(&models.AnalystRating{}); err != nil {
		log.Fatal("Table migration error: ", err)
	}
}
