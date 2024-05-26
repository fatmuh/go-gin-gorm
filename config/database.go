package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func DatabaseConnection() *gorm.DB {
	dsn := "root@tcp(localhost:3306)/ecom?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	return db
}
