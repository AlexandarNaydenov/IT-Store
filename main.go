package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.tools.sap/I546075/empty_project/pkg/config"
	"github.tools.sap/I546075/empty_project/pkg/database"
	"github.tools.sap/I546075/empty_project/pkg/handlers"
	"github.tools.sap/I546075/empty_project/pkg/services"
)

func main() {
	config.InitConfig()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	database, err := database.NewMongoDB(ctx)
	if err != nil {
		slog.Error("Failed to initialize the database: %v", err)
		os.Exit(1)
	}

	controller := services.NewProductService(database)
	handler := handlers.NewProductHandler(controller)

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		products := v1.Group("/products")
		{
			products.GET("", handler.GetProducts)
			products.GET(":id", handler.GetProductByID)
			products.POST("", handler.CreateProduct)
			// products.PUT("/:id", handler.UpdateProductByID)
		}
	}

	if err := r.Run(config.Config().Server.Port); err != nil {
		slog.Error("Server stopped: %v", err)
		os.Exit(1)
	}
}
