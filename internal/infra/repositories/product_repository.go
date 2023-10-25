package repositories

import (
	"errors"
	"fmt"
	"g37-lanchonete/internal/domain/models"
	"g37-lanchonete/internal/domain/ports"
	"g37-lanchonete/internal/domain/services/dto"
	"g37-lanchonete/internal/infra/clients"
	"strconv"
)

type productRepository struct {
	client clients.SQLClient
}

func NewProductRepository(client clients.SQLClient) ports.ProductRepository {
	return productRepository{
		client: client,
	}
}

func (r productRepository) FindAllProducts(pageParams dto.PageParams) ([]models.Product, error) {
	var products []models.Product
	err := r.client.FindAll(&products, pageParams.GetLimit(), pageParams.GetOffset())
	if err != nil {
		return nil, fmt.Errorf("failed to find all products, error %v", err)
	}

	return products, nil
}

func (r productRepository) FindProductsByCategory(pageParams dto.PageParams, category string) ([]models.Product, error) {
	var products []models.Product
	err := r.client.Find(&products, pageParams.GetLimit(), pageParams.GetOffset(), "category = ?", category)
	if err != nil {
		return nil, fmt.Errorf("failed to find products by id, error %v", err)
	}

	return products, nil
}

func (r productRepository) SaveProduct(product models.Product) error {
	err := r.client.Save(&product)
	if err != nil {
		return fmt.Errorf("failed to save product, error %v", err)
	}

	return nil
}

func (r productRepository) UpdateProduct(id uint, product models.Product) error {
	var oldProduct models.Product
	err := r.client.FindById(strconv.FormatUint(uint64(id), 10), &oldProduct)
	if err != nil {
		if errors.Is(err, clients.ErrNotFound) {
			return fmt.Errorf("product [%d] not found, error %v", id, err)
		}
		return fmt.Errorf("failed to find the product [%d], error %v", id, err)
	}

	product.ID = id
	err = r.client.Save(&product)
	if err != nil {
		return fmt.Errorf("failed to update the product, error %v", err)
	}

	return nil
}

func (r productRepository) DeleteProduct(id uint) error {
	var product models.Product
	product.ID = id

	err := r.client.Delete(&product)
	if err != nil {
		return fmt.Errorf("failed to delete the product [%d], error %v", id, err)
	}

	return nil
}
