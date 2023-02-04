package usecase

import "github.com/rafaelsous/fullcycle12-go-esquenta/internal/entity"

type ListProductsOutputDTO struct {
	ID    string
	Name  string
	Price float64
}

type ListProductsUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewListProductsUseCase(productRepository entity.ProductRepository) *ListProductsUseCase {
	return &ListProductsUseCase{ProductRepository: productRepository}
}

func (u *ListProductsUseCase) Execute() ([]*ListProductsOutputDTO, error) {
	products, err := u.ProductRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var productsOuput []*ListProductsOutputDTO
	for _, product := range products {
		productsOuput = append(productsOuput, &ListProductsOutputDTO{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}

	return productsOuput, nil
}
