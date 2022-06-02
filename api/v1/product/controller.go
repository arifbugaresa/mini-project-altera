package product

import (
	"github.com/arifbugaresa/mini-project-altera/api/v1/common"
	"github.com/arifbugaresa/mini-project-altera/api/v1/product/request"
	"github.com/arifbugaresa/mini-project-altera/business/product"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service product.Service
}

func NewController(service product.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) InsertProduct(context echo.Context) error {
	var insertProductRequest request.InsertProductRequest

	if err := context.Bind(&insertProductRequest); err != nil {
	}

	err := controller.service.InsertProduct(insertProductRequest)
	if err != nil {
		return context.JSON(common.NewErrorBusinessResponse(err))
	}

	return context.JSON(common.NewSuccessResponseWithoutData())
}
