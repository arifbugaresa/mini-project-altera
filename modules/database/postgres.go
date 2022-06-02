package database

import (
	"fmt"
	"github.com/arifbugaresa/mini-project-altera/config"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

func NewDatabaseConnection(config *config.AppConfig) *gorm.DB {
	var dsn string

	log.Info("Starting on " + config.AppEnvironment)
	if config.AppEnvironment == "development" {
		dsn = fmt.Sprintf(`host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta`,
			config.DBHost, config.DBUsername, config.DBPassword, strconv.Itoa(config.DBPort), config.DBName)
	} else if config.AppEnvironment == "sandbox" {
		dsn = fmt.Sprintf(`host=%s user=%s password=%s port=%s dbname=%s sslmode=require TimeZone=Asia/Jakarta`,
			config.DBHost, config.DBUsername, config.DBPassword, strconv.Itoa(config.DBPort), config.DBName)
	}

	dbGormPostgres, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Failed")
	} else {
		log.Info("DB Connection Success")
	}

	return dbGormPostgres
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()

	log.Info("DB Connection Close")
}