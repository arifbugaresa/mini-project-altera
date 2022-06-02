package main

import (
	"fmt"
	"github.com/arifbugaresa/mini-project-altera/api"
	productController "github.com/arifbugaresa/mini-project-altera/api/v1/product"
	product "github.com/arifbugaresa/mini-project-altera/business/product"
	productService "github.com/arifbugaresa/mini-project-altera/business/product"
	config "github.com/arifbugaresa/mini-project-altera/config"
	"github.com/arifbugaresa/mini-project-altera/modules/database"
	productRepository "github.com/arifbugaresa/mini-project-altera/modules/product"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"strconv"
)

func main() {
	// load configuration
	config := config.GetConfig()

	// initialize database connection
	dbGormPostgres := database.NewDatabaseConnection(config)
	defer database.CloseDatabaseConnection(dbGormPostgres)

	// migration
	dbGormPostgres.AutoMigrate(&product.Product{})

	// initialize product
	productRepo := productRepository.NewRepository(dbGormPostgres)
	productService := productService.NewService(productRepo)
	productController := productController.NewController(productService)


	// create echo http
	e := echo.New()
	api.APIController(e, productController)

	// run server
	address := fmt.Sprintf("localhost:%s", strconv.Itoa(config.AppPort))
	err := e.Start(address)
	if err != nil {
		log.Info("Shutting down the server")
	}

}
