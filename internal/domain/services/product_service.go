package services

import (
	"g37-lanchonete/internal/domain/models"
	"g37-lanchonete/internal/domain/ports"
	"g37-lanchonete/internal/domain/services/dto"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type productService struct {
	productRepository ports.ProductRepository
}

func NewProductService(productRepository ports.ProductRepository) ports.ProductService {
	return productService{
		productRepository: productRepository,
	}
}

func (s productService) GetAllProducts(pageParameters dto.PageParams) (dto.Page[models.Product], error) {
	products, err := s.productRepository.FindAllProducts(pageParameters)
	if err != nil {
		log.Errorf("failed to get all products, error: %v", err)
		return dto.Page[models.Product]{}, err
	}

	page := dto.BuildPage[models.Product](products, pageParameters)
	return page, nil
}

func (s productService) GetProductsByCategory(pageParameters dto.PageParams, category string) (dto.Page[models.Product], error) {
	products, err := s.productRepository.FindProductsByCategory(pageParameters, category)
	if err != nil {
		log.Errorf("failed to get products by category, error: %v", err)
		return dto.Page[models.Product]{}, err
	}

	page := dto.BuildPage[models.Product](products, pageParameters)
	return page, nil
}

func (s productService) CreateProduct(productDTO dto.ProductDTO) error {
	product := productDTO.ToProduct()

	err := s.productRepository.SaveProduct(product)
	if err != nil {
		log.Errorf("failed to save product, error: %v", err)
		return err
	}

	return nil
}

func (s productService) UpdateProduct(idStr string, productDTO dto.ProductDTO) error {
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.Errorf("failed to parse id [%s], error: %v", idStr, err)
		return err
	}

	product := productDTO.ToProduct()
	err = s.productRepository.UpdateProduct(uint(id), product)
	if err != nil {
		log.Errorf("failed to update product, error: %v", err)
		return err
	}

	return nil
}

func (s productService) DeleteProduct(idStr string) error {
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.Errorf("failed to parse id [%s], error: %v", idStr, err)
		return err
	}

	err = s.productRepository.DeleteProduct(uint(id))
	if err != nil {
		log.Errorf("failed to delete product, error: %v", err)
		return err
	}

	return nil
}
