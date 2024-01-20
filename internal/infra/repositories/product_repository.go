package repositories

import (
	"fmt"
	"g37-lanchonete/internal/core/domain"
	"g37-lanchonete/internal/core/ports"
	"g37-lanchonete/internal/core/services/dto"
	"g37-lanchonete/internal/infra/clients/sql"
	"g37-lanchonete/internal/infra/sqlscripts"
)

type productRepository struct {
	sqlClient sql.SQLClient
}

func NewProductRepository(sqlClient sql.SQLClient) ports.ProductRepository {
	return productRepository{
		sqlClient: sqlClient,
	}
}

func (r productRepository) FindAllProducts(pageParams dto.PageParams) ([]domain.Product, error) {
	getAllProductsQuery := fmt.Sprintf(sqlscripts.GetAllProductsQuery, pageParams.GetLimit(), pageParams.GetOffset())

	rows, err := r.sqlClient.Find(getAllProductsQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to find all products, error %w", err)
	}
	defer rows.Close()

	products := []domain.Product{}
	for rows.Next() {
		var product domain.Product
		err = rows.Scan(&product.ID, &product.Name, &product.SkuId, &product.Description, &product.Category, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan all products, error %w", err)
		}

		products = append(products, product)
	}

	return products, nil
}

func (r productRepository) FindProductsByCategory(pageParams dto.PageParams, category string) ([]domain.Product, error) {
	getProductsByCategoryQuery := fmt.Sprintf(sqlscripts.GetProductsByCategoryQuery, pageParams.GetLimit(), pageParams.GetOffset())

	rows, err := r.sqlClient.Find(getProductsByCategoryQuery, category)
	if err != nil {
		return nil, fmt.Errorf("failed to find products by category, error %w", err)
	}
	defer rows.Close()

	products := []domain.Product{}
	for rows.Next() {
		var product domain.Product
		err = rows.Scan(&product.ID, &product.Name, &product.SkuId, &product.Description, &product.Category, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan products by category, error %w", err)
		}

		products = append(products, product)
	}

	return products, nil
}

func (r productRepository) FindProductById(id int) (domain.Product, error) {
	row := r.sqlClient.FindOne(sqlscripts.GetProductByIdQuery, id)

	var product domain.Product
	err := row.Scan(&product.ID, &product.Name, &product.SkuId, &product.Description, &product.Category, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return domain.Product{}, fmt.Errorf("failed to find product by id, error %w", err)
	}

	return product, nil
}

func (r productRepository) SaveProduct(product domain.Product) error {
	inserProductCmd := fmt.Sprintf(sqlscripts.InsertProductCmd)

	_, err := r.sqlClient.Exec(inserProductCmd, product.Name, product.SkuId, product.Description, product.Category,
		product.Price, product.CreatedAt, product.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to save product, error %w", err)
	}

	return nil
}

func (r productRepository) UpdateProduct(id int, product domain.Product) error {
	updateProductCmd := fmt.Sprintf(sqlscripts.UpdateProductCmd)

	result, err := r.sqlClient.Exec(updateProductCmd, id, product.Name, product.SkuId, product.Description, product.Category,
		product.Price, product.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to update the product [%d], error %w", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected on updating product [%d], error %w", id, err)
	}

	if rowsAffected < 1 {
		return sql.ErrNotFound
	}

	return nil
}

func (r productRepository) DeleteProduct(id int) error {
	deleteProductCmd := fmt.Sprintf(sqlscripts.DeleteProductCmd)

	result, err := r.sqlClient.Exec(deleteProductCmd, id)
	if err != nil {
		return fmt.Errorf("failed to delete the product [%d], error %v", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected on deleting product [%d], error %w", id, err)
	}

	if rowsAffected < 1 {
		return sql.ErrNotFound
	}
	return nil
}
