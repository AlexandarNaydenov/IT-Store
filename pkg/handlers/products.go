package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.tools.sap/I546075/empty_project/pkg/models"
)

func NewProductHandler(controller ProductService) *productHandler {
	return &productHandler{
		controller: controller,
	}
}

type ProductService interface {
	GetProducts(ctx context.Context) ([]models.Product, error)
	GetProductByID(ctx context.Context, id string) (*models.Product, error)
	CreateProduct(ctx context.Context, product models.Product) (string, error)
}

type productHandler struct {
	controller ProductService
}

// GetProducts send a json response with all the products to the client
func (h *productHandler) GetProducts(ctx *gin.Context) {
	products, err := h.controller.GetProducts(ctx)
	if err != nil {
		slog.Error("Failed to fetch the products: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

// GetProductByID send a json response with the product to the client
func (h *productHandler) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := h.controller.GetProductByID(ctx, id)
	if err != nil {
		slog.Error("Failed to fetch the product with id %s: %v", id, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// CreateProduct creates a new product and sends a json response with the id to the client
func (h *productHandler) CreateProduct(ctx *gin.Context) {
	var product models.Product
	if err := ctx.BindJSON(&product); err != nil {
		slog.Error("Failed to bind the product: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.controller.CreateProduct(ctx, product)
	if err != nil {
		slog.Error("Failed to create the product: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}
