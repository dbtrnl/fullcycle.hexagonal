package application

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {	return nil, err }

	return product, nil
}

func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price

	// If product is invalid
	_, err := product.IsValid()
	if err != nil {	return &Product{}, err }

	// Save product on DB
	result, err := s.Persistence.Save(product)
	if err != nil {	return &Product{}, err }

	return result, nil
}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	// Enable the product
	err := product.Enable()
	if err != nil {	return &Product{}, err }

	// Save the enabled product on DB
	result, err := s.Persistence.Save(product)
	if err != nil {	return nil, err }

	return result, nil
}

func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	// Disable the product
	err := product.Disable()
	if err != nil {	return &Product{}, err }

	// Save the disabled product on DB
	result, err := s.Persistence.Save(product)
	if err != nil {	return nil, err }

	return result, nil
}
