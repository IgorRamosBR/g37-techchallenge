package repositories

import (
	"errors"
	"fmt"
	"g37-lanchonete/internal/core/domain"
	"g37-lanchonete/internal/core/ports"
	"g37-lanchonete/internal/core/services/dto"
	"g37-lanchonete/internal/infra/clients"
)

type productRepository struct {
	client clients.SQLClient
}

func NewProductRepository(client clients.SQLClient) ports.ProductRepository {
	return productRepository{
		client: client,
	}
}

func (r productRepository) FindAllProducts(pageParams dto.PageParams) ([]domain.Product, error) {
	var products []domain.Product
	err := r.client.FindAll(&products, pageParams.GetLimit(), pageParams.GetOffset(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to find all products, error %v", err)
	}

	return products, nil
}

func (r productRepository) FindProductsByCategory(pageParams dto.PageParams, category string) ([]domain.Product, error) {
	var products []domain.Product
	err := r.client.Find(&products, pageParams.GetLimit(), pageParams.GetOffset(), "category = ?", category)
	if err != nil {
		return nil, fmt.Errorf("failed to find products by category, error %v", err)
	}

	return products, nil
}

func (r productRepository) FindProductById(id int) (domain.Product, error) {
	var product domain.Product
	err := r.client.FindById(id, &product)
	if err != nil {
		return domain.Product{}, fmt.Errorf("failed to find product by id, error %v", err)
	}

	return product, nil
}

func (r productRepository) SaveProduct(product domain.Product) error {
	err := r.client.Save(&product)
	if err != nil {
		return fmt.Errorf("failed to save product, error %v", err)
	}

	return nil
}

func (r productRepository) UpdateProduct(id int, product domain.Product) error {
	var oldProduct domain.Product
	err := r.client.FindById(int(id), &oldProduct)
	if err != nil {
		if errors.Is(err, clients.ErrNotFound) {
			return fmt.Errorf("product [%d] not found, error %v", id, err)
		}
		return fmt.Errorf("failed to find the product [%d], error %v", id, err)
	}

	product.ID = id
	err = r.client.Update(&product)
	if err != nil {
		return fmt.Errorf("failed to update the product, error %v", err)
	}

	return nil
}

func (r productRepository) DeleteProduct(id int) error {
	var product domain.Product
	product.ID = id

	err := r.client.Delete(&product)
	if err != nil {
		return fmt.Errorf("failed to delete the product [%d], error %v", id, err)
	}

	return nil
}
