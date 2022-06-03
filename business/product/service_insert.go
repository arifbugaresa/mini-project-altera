package product

import (
	"github.com/arifbugaresa/mini-project-altera/api/v1/product/request"
	"github.com/arifbugaresa/mini-project-altera/business"
)

// InsertProduct is used for add new product
func (s *service) InsertProduct(userParam request.InsertProductRequest) error{

	// validate and read struct input product
	product, err := s.readBodyAndValidate(userParam)
	if err != nil {
		return err
	}

	// insert into database
	err = s.repository.InsertProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func (s service) readBodyAndValidate(userParam request.InsertProductRequest) (product Product, err error) {

	// validate input struct
	err = userParam.MandatoryValidation()
	if err != nil {
		return product, business.ErrInvalidBody
	}

	// read, convert struct dto to model
	product = convertDTOToModel(userParam)

	return product,nil
}

func convertDTOToModel(userParam request.InsertProductRequest) (product Product) {
	return Product{
		Name:       userParam.Name,
		Price:      userParam.Price,
		CreatedBy:  userParam.CreatedBy,
	}
}