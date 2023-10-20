package application

import (
	"errors"
	"g37-lanchonete/internal/domain/ports"
	"g37-lanchonete/internal/domain/services/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService ports.ProductService
}

func NewProductHandler(productService ports.ProductService) ProductHandler {
	return ProductHandler{
		productService: productService,
	}
}

func (h ProductHandler) GetProducts(c *gin.Context) {
	category := c.Query("category")
	if category != "" {
		h.getProductsByCategory(c, category)
		return
	}
	h.getAllProducts(c)
}

func (h ProductHandler) CreateProducts(c *gin.Context) {
	var product dto.ProductDTO
	err := c.ShouldBindJSON(&product)
	if err != nil {
		handleBadRequestResponse(c, "failed to bind product payload", err)
		return
	}

	valid, err := product.ValidateProduct()
	if !valid {
		handleBadRequestResponse(c, "invalid product payload", err)
		return
	}

	err = h.productService.CreateProduct(product)
	if err != nil {
		handleInternalServerResponse(c, "failed to create product", err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		handleBadRequestResponse(c, "id path param is required", errors.New("id path parameter is missing"))
		return
	}

	var product dto.ProductDTO
	err := c.ShouldBindJSON(&product)
	if err != nil {
		handleBadRequestResponse(c, "failed to bind product payload", err)
		return
	}

	valid, err := product.ValidateProduct()
	if !valid {
		handleBadRequestResponse(c, "invalid product payload", err)
		return
	}

	err = h.productService.UpdateProduct(id, product)
	if err != nil {
		handleInternalServerResponse(c, "failed to create product", err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		handleBadRequestResponse(c, "id path param is required", errors.New("id path parameter is missing"))
		return
	}

	err := h.productService.DeleteProduct(id)
	if err != nil {
		handleInternalServerResponse(c, "failed to delete product", err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h ProductHandler) getAllProducts(c *gin.Context) {
	products, err := h.productService.GetAllProducts()
	if err != nil {
		handleInternalServerResponse(c, "failed to get all products", err)
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h ProductHandler) getProductsByCategory(c *gin.Context, category string) {
	products, err := h.productService.GetProductsByCategory(category)
	if err != nil {
		handleInternalServerResponse(c, "failed to get products by category", err)
		return
	}
	c.JSON(http.StatusOK, products)
}
