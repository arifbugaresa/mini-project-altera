package request

import "github.com/arifbugaresa/mini-project-altera/business"

/*
Data Transfer Object (IN)
*/

type InsertProductRequest struct {
	Name      string `json:"name"`
	Price     uint   `json:"price"`
	CreatedBy string `json:"created_by"`
}

func (pr InsertProductRequest) MandatoryValidation() error {
	if pr.Name == "" {
		return business.ErrInvalidBody
	}

	if pr.Price < 1 {
		return business.ErrInvalidBody
	}

	return nil
}