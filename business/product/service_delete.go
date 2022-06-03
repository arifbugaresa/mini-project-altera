package product

func (s *service) DeleteProductByID(id int) error {
	return s.repository.DeleteProductByID(id)
}
