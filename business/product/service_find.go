package product

func (s *service) FindProductByID(id int) (Product, error) {
	return s.repository.FindProductByID(id)
}

func (s *service) FindAllProduct() (Database, error) {
	var (
		listProduct []Product
		database Database
	)

	listProduct, err := s.repository.FindAllProduct()
	if err != nil {
		return Database{}, err
	}

	database.Products = listProduct

	return database, nil
}