package repository

import (
	"Chopie/internal/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		panic(err)
	}
	return db
}
