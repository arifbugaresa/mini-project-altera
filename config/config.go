package config

import (
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"os"
	"strconv"
	"sync"
)

type AppConfig struct {
	AppPort        int
	AppEnvironment string
	DBDriver       string
	DBHost         string
	DBPort         int
	DBUsername     string
	DBPassword     string
	DBName         string
}

var (
	appConfig *AppConfig
	lock      = &sync.Mutex{}
)

// GetConfig Initialize config in singleton way
func GetConfig() *AppConfig {
	if appConfig != nil {
		return appConfig
	}

	// Untuk mengunci perubahan yg terjadi bersamaan, sesuai konsep singleton
	lock.Lock()
	defer lock.Unlock()

	// re-check after locking
	if appConfig != nil {
		return appConfig
	}

	appConfig = initConfig()

	return appConfig
}

// initConfig read configuration from .env file
func initConfig() *AppConfig {
	var appConfig AppConfig

	err := godotenv.Load("config/.env")
	if err != nil {
		log.Infof("failed load file environment")
	} else {
		log.Infof("success read file environment")
	}

	// Set Environtment
	appPort, _ := strconv.Atoi(os.Getenv("APP_PORT"))
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	appConfig.AppPort = appPort
	appConfig.AppEnvironment = os.Getenv("APP_ENV")

	appConfig.DBDriver = os.Getenv("DB_DRIVER")
	appConfig.DBPort = dbPort
	appConfig.DBHost = os.Getenv("DB_HOST")
	appConfig.DBUsername = os.Getenv("DB_USERNAME")
	appConfig.DBPassword = os.Getenv("DB_PASSWORD")
	appConfig.DBName = os.Getenv("DB_NAME")

	return &appConfig
}
