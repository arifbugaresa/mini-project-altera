package api

import (
	"github.com/arifbugaresa/mini-project-altera/api/v1/product"
	echo "github.com/labstack/echo/v4"
)

// APIController register all v1 API with routing path
func APIController(e *echo.Echo, productController *product.Controller)  {

	// Helper Health
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})

	// Product
	productV1 := e.Group("v1/products")
	productV1.POST("", productController.InsertProduct)

}