package config

import (
	"log"
	"stock-trades-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "host=localhost user=postgres password=your_passowrd dbname=stock_trades_db port=5432 sslmode=disable"
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatal("Failed to connect to the database", err)
    }

    database.AutoMigrate(&models.User{}, &models.Trade{})

    DB = database
}

