package services

import (
	"context"
	"fmt"

	"github.tools.sap/I546075/empty_project/pkg/models"
)

func NewProductService(storage ProductsStorage) *productService {
	return &productService{
		storage: storage,
	}
}

type ProductsStorage interface {
	GetProducts(ctx context.Context) ([]models.Product, error)
	GetProductByID(ctx context.Context, id string) (*models.Product, error)
	CreateProduct(ctx context.Context, product models.Product) (string, error)
}

type productService struct {
	storage ProductsStorage
}

// GetProducts returns all the products from the storage
func (c *productService) GetProducts(ctx context.Context) ([]models.Product, error) {
	products, err := c.storage.GetProducts(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get the products: %w", err)
	}

	return products, nil
}

// GetProductByID returns the product with the given id from the storage
func (c *productService) GetProductByID(ctx context.Context, id string) (*models.Product, error) {
	product, err := c.storage.GetProductByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get the product with id %s: %w", id, err)
	}

	return product, nil
}

// CreateProduct creates a new product in the storage and returns the id
func (c *productService) CreateProduct(ctx context.Context, product models.Product) (string, error) {
	id, err := c.storage.CreateProduct(ctx, product)
	if err != nil {
		return "", fmt.Errorf("failed to create the product: %w", err)
	}

	return id, nil
}
