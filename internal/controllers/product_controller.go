package controllers

import (
	"errors"
	"g37-lanchonete/internal/core/ports"
	"g37-lanchonete/internal/core/services/dto"
	"g37-lanchonete/internal/infra/drivers/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService ports.ProductService
}

func NewProductController(productService ports.ProductService) ProductController {
	return ProductController{
		productService: productService,
	}
}

func (c ProductController) GetProducts(ctx *gin.Context) {
	category := ctx.Query("category")
	pageParams, err := getPageParams(ctx)
	if err != nil {
		handleBadRequestResponse(ctx, "invalid query parameters", err)
	}

	if category != "" {
		c.getProductsByCategory(ctx, pageParams, category)
		return
	}

	c.getAllProducts(ctx, pageParams)
}

func (c ProductController) CreateProducts(ctx *gin.Context) {
	var product dto.ProductDTO
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		handleBadRequestResponse(ctx, "failed to bind product payload", err)
		return
	}

	valid, err := product.ValidateProduct()
	if !valid {
		handleBadRequestResponse(ctx, "invalid product payload", err)
		return
	}

	err = c.productService.CreateProduct(product)
	if err != nil {
		handleInternalServerResponse(ctx, "failed to create product", err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c ProductController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		handleBadRequestResponse(ctx, "id path param is required", errors.New("id path parameter is missing"))
		return
	}

	var product dto.ProductDTO
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		handleBadRequestResponse(ctx, "failed to bind product payload", err)
		return
	}

	valid, err := product.ValidateProduct()
	if !valid {
		handleBadRequestResponse(ctx, "invalid product payload", err)
		return
	}

	err = c.productService.UpdateProduct(id, product)
	if err != nil {
		if errors.Is(err, sql.ErrNotFound) {
			handleNotFoundResponse(ctx, "product not found", err)
			return
		}
		handleInternalServerResponse(ctx, "failed to create product", err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c ProductController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		handleBadRequestResponse(ctx, "id path param is required", errors.New("id path parameter is missing"))
		return
	}

	err := c.productService.DeleteProduct(id)
	if err != nil {
		if errors.Is(err, sql.ErrNotFound) {
			handleNotFoundResponse(ctx, "product not found", err)
			return
		}
		handleInternalServerResponse(ctx, "failed to delete product", err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c ProductController) getAllProducts(ctx *gin.Context, pageParameters dto.PageParams) {
	products, err := c.productService.GetAllProducts(pageParameters)
	if err != nil {
		handleInternalServerResponse(ctx, "failed to get all products", err)
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (c ProductController) getProductsByCategory(ctx *gin.Context, pageParameters dto.PageParams, category string) {
	products, err := c.productService.GetProductsByCategory(pageParameters, category)
	if err != nil {
		handleInternalServerResponse(ctx, "failed to get products by category", err)
		return
	}
	ctx.JSON(http.StatusOK, products)
}
