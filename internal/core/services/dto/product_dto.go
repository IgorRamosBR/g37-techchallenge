package dto

import (
	"g37-lanchonete/internal/core/entities"

	"github.com/asaskevich/govalidator"
)

type ProductDTO struct {
	Name        string  `json:"name" valid:"length(0|100)~Name length should be less than 100 characters"`
	SkuId       string  `json:"skuId" valid:"length(0|50)~Sku length should be less than 50 characters"`
	Description string  `json:"description" valid:"length(0|2000)~Description length should be less than 2000 characters"`
	Category    string  `json:"category" valid:"length(0|60)~Category length should be less than 60 characters"`
	Price       float64 `json:"price" valid:"float,required~Price is required|range(0.01|)~Price greater than 0.00"`
}

func (p ProductDTO) ToProduct() entities.Product {
	return entities.Product{
		Name:        p.Name,
		SkuId:       p.SkuId,
		Description: p.Description,
		Category:    p.Category,
		Price:       p.Price,
	}
}

func (p ProductDTO) ValidateProduct() (bool, error) {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return false, err
	}

	return true, nil
}
