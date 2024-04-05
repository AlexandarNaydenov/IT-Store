package database

import (
	"context"
	"fmt"

	"github.tools.sap/I546075/empty_project/pkg/config"
	"github.tools.sap/I546075/empty_project/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB(ctx context.Context) (*mongodb, error) {
	client, err := initClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize the database client: %w", err)
	}

	return &mongodb{
		client:             client,
		databaseName:       config.Config().Database.Name,
		productsCollection: config.Config().Database.Products.Collection,
	}, nil
}

type mongodb struct {
	client             *mongo.Client
	databaseName       string
	productsCollection string
}

// GetProducts fetches all the products from the database
func (db *mongodb) GetProducts(ctx context.Context) (products []models.Product, err error) {
	collection := db.client.Database(db.databaseName).Collection(db.productsCollection)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the products: %w", err)
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &products); err != nil {
		return nil, fmt.Errorf("failed to decode the products: %w", err)
	}

	return products, nil
}

// GetProductByID fetches the product with the given id from the database
func (db *mongodb) GetProductByID(ctx context.Context, id string) (product *models.Product, err error) {
	collection := db.client.Database(db.databaseName).Collection(db.productsCollection)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("failed to convert the id %s to an ObjectID: %w", id, err)
	}

	filter := bson.D{{"_id", objID}}
	if err = collection.FindOne(ctx, filter).Decode(&product); err != nil {
		return nil, fmt.Errorf("failed to fetch the product with id %s: %w", id, err)
	}

	return product, nil
}

// CreateProduct inserts a new product into the database
func (db *mongodb) CreateProduct(ctx context.Context, product models.Product) (string, error) {
	collection := db.client.Database(db.databaseName).Collection(db.productsCollection)

	request := bson.M{"name": product.Name, "price": product.Price}
	result, err := collection.InsertOne(ctx, request)
	if err != nil {
		return "", fmt.Errorf("failed to insert the product: %w", err)
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func initClient(ctx context.Context) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", config.Config().Database.Username, config.Config().Database.Password, config.Config().Database.Host, config.Config().Database.Port)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	return client, nil
}
