package services

import (
	"g37-lanchonete/internal/core/domain"
	"g37-lanchonete/internal/core/ports"
	"g37-lanchonete/internal/core/services/dto"
	"g37-lanchonete/internal/infra/gateways"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

type productService struct {
	productRepositoryGateway gateways.ProductRepositoryGateway
}

func NewProductService(productRepositoryGateway gateways.ProductRepositoryGateway) ports.ProductService {
	return productService{
		productRepositoryGateway: productRepositoryGateway,
	}
}

func (s productService) GetAllProducts(pageParameters dto.PageParams) (dto.Page[domain.Product], error) {
	products, err := s.productRepositoryGateway.FindAllProducts(pageParameters)
	if err != nil {
		log.Errorf("failed to get all products, error: %v", err)
		return dto.Page[domain.Product]{}, err
	}

	page := dto.BuildPage[domain.Product](products, pageParameters)
	return page, nil
}

func (s productService) GetProductsByCategory(pageParameters dto.PageParams, category string) (dto.Page[domain.Product], error) {
	products, err := s.productRepositoryGateway.FindProductsByCategory(pageParameters, category)
	if err != nil {
		log.Errorf("failed to get products by category, error: %v", err)
		return dto.Page[domain.Product]{}, err
	}

	page := dto.BuildPage[domain.Product](products, pageParameters)
	return page, nil
}

func (s productService) GetProductById(id int) (domain.Product, error) {
	product, err := s.productRepositoryGateway.FindProductById(id)
	if err != nil {
		log.Errorf("failed to get product by id, error: %v", err)
		return domain.Product{}, err
	}

	return product, nil
}

func (s productService) CreateProduct(productDTO dto.ProductDTO) error {
	product := productDTO.ToProduct()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	err := s.productRepositoryGateway.SaveProduct(product)
	if err != nil {
		log.Errorf("failed to save product, error: %v", err)
		return err
	}

	return nil
}

func (s productService) UpdateProduct(idStr string, productDTO dto.ProductDTO) error {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Errorf("failed to parse id [%s], error: %v", idStr, err)
		return err
	}

	product := productDTO.ToProduct()
	product.UpdatedAt = time.Now()
	err = s.productRepositoryGateway.UpdateProduct(id, product)
	if err != nil {
		log.Errorf("failed to update product, error: %v", err)
		return err
	}

	return nil
}

func (s productService) DeleteProduct(idStr string) error {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Errorf("failed to parse id [%s], error: %v", idStr, err)
		return err
	}

	err = s.productRepositoryGateway.DeleteProduct(id)
	if err != nil {
		log.Errorf("failed to delete product, error: %v", err)
		return err
	}

	return nil
}
