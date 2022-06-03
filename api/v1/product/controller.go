package product

import (
	"github.com/arifbugaresa/mini-project-altera/api/v1/common"
	"github.com/arifbugaresa/mini-project-altera/api/v1/product/request"
	"github.com/arifbugaresa/mini-project-altera/business/product"
	"github.com/labstack/echo/v4"
	"strconv"
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

func (controller *Controller) FindProductByID(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return context.JSON(common.NewBadRequestResponse())
	}

	product, err := controller.service.FindProductByID(id)
	if err != nil {
		return context.JSON(common.NewErrorBusinessResponse(err))
	}

	return context.JSON(common.NewSuccessResponse(product))
}

func (controller *Controller) FindAllProduct(context echo.Context) error {
	listProduct, err := controller.service.FindAllProduct()
	if err != nil {
		return context.JSON(common.NewErrorBusinessResponse(err))
	}

	return context.JSON(common.NewSuccessResponse(listProduct))
}

func (controller *Controller) DeleteProductByID(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return context.JSON(common.NewBadRequestResponse())
	}

	// find product by id
	product, err := controller.service.FindProductByID(id)
	if err != nil {
		return context.JSON(common.NewErrorBusinessResponse(err))
	}

	// delete product
	err = controller.service.DeleteProductByID(int(product.ID))
	if err != nil {
		return context.JSON(common.NewErrorBusinessResponse(err))
	}

	return context.JSON(common.NewSuccessResponseWithoutData())
}
