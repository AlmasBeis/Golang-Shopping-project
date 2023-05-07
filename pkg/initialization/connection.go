package initialization

import (
	"final-project/pkg/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB(cfg *Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User,
		cfg.Password, cfg.DBName, cfg.DBPort, cfg.SSLMode)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	err = DB.AutoMigrate(&models.Category{}, &models.Item{}, &models.Comment{}, &models.Purchase{}, &models.User{})
	if err != nil {
		log.Fatal("Failed to migrate")
	}
	fmt.Println("? Connected successfully to the Database")

}
