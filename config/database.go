package config

import (
	"fmt"
	"log"
	"os"

	"github.com/VieiraGabrielAlexandre/luztecnologia-cms-backend/core/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DB_USERNAME") +
		":" +
		os.Getenv("DB_PASSWORD") +
		"@tcp(" +
		os.Getenv("DB_HOST") +
		":" +
		os.Getenv("DB_PORT") +
		")/" +
		os.Getenv("DB_DATABASE") +
		"?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Executing migrations ...")

	db.AutoMigrate(&models.Client{})

	fmt.Println("Migrations executed successfully")

	fmt.Println("Sucess")

	DB = db
}
