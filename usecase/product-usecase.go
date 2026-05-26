package usecase

import (
	"api-go/model"
	"api-go/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repository repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repository,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUseCase) GetProductById(id int) (*model.Product, error) {

	product, err := pu.repository.GetProductById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pu *ProductUseCase) DeleteProduct(id int) error {
	err := pu.repository.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}

func (pu *ProductUseCase) UpdateProduct(product model.Product) (model.Product, error) {
	updatedProduct, err := pu.repository.UpdateProduct(product)
	if err != nil {
		return model.Product{}, err
	}
	return updatedProduct, nil
}
