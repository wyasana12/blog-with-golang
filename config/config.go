package config

import (
	"blog-go/internal/model"
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Validate *validator.Validate

func ConnectDB() {
	_ = godotenv.Load(".env")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail Connect To DB:", err)
	}

	DB = db

	DB.AutoMigrate(&model.User{}, &model.Role{})
	fmt.Println("Success Connect to Database")
}

func Init() {
	Validate = validator.New()
	log.Println("Validator Running Success...")
}

func SeedRoles() {
	roles := []string{"admin"}

	for _, role := range roles {
		DB.FirstOrCreate(&model.Role{}, model.Role{Name: role})
	}
}
