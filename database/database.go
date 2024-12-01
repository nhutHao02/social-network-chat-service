package database

import (
	"context"
	"time"

	"github.com/nhutHao02/social-network-chat-service/config"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type MongoDbClient struct {
	Mdb *mongo.Client
}

func ConnectToMongo(cfg *config.DatabaseConfig) *MongoDbClient {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.ConnectionString))
	if err != nil {
		logger.Fatal("Failed to connect to MongoDB", zap.Error(err))
	}

	// Test the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Fatal("Failed to ping MongoDB", zap.Error(err))
	}

	logger.Info("Connected to MongoDB successfully!!!")
	return &MongoDbClient{
		Mdb: client,
	}
}

// FindOne fetches a single document from the specified collection.
func (db *MongoDbClient) FindOne(ctx context.Context, databaseName, collectionName string, result interface{}, filter interface{}) error {
	collection := db.Mdb.Database(databaseName).Collection(collectionName)
	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

// FindMany fetches multiple documents from the specified collection.
func (db *MongoDbClient) FindMany(ctx context.Context, databaseName, collectionName string, filter interface{}, result interface{}, opts ...*options.FindOptions) error {
	collection := db.Mdb.Database(databaseName).Collection(collectionName)
	cursor, err := collection.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	// Decode the result into the provided slice
	if err := cursor.All(ctx, result); err != nil {
		return err
	}

	return nil
}

// InsertOne inserts a single document into the specified collection.
func (db *MongoDbClient) InsertOne(ctx context.Context, databaseName, collectionName string, document interface{}) (interface{}, error) {
	collection := db.Mdb.Database(databaseName).Collection(collectionName)
	result, err := collection.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

// InsertMany inserts documents into the specified collection.
func (db *MongoDbClient) InsertMany(ctx context.Context, databaseName, collectionName string, documents []interface{}) ([]interface{}, error) {
	collection := db.Mdb.Database(databaseName).Collection(collectionName)
	result, err := collection.InsertMany(ctx, documents)
	if err != nil {
		return nil, err
	}
	return result.InsertedIDs, nil
}

// UpdateOne updates a single document in the specified collection based on the filter.
func (db *MongoDbClient) UpdateOne(ctx context.Context, databaseName, collectionName string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	collection := db.Mdb.Database(databaseName).Collection(collectionName)
	updateResult, err := collection.UpdateOne(ctx, filter, update, opts...)
	if err != nil {
		return nil, err
	}
	return updateResult, nil
}

// UpdateMany updates documents in the specified collection based on the filter.
func (db *MongoDbClient) UpdateMany(ctx context.Context, databaseName, collectionName string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	collection := db.Mdb.Database(databaseName).Collection(collectionName)
	updateResult, err := collection.UpdateMany(ctx, filter, update, opts...)
	if err != nil {
		return nil, err
	}
	return updateResult, nil
}

// DeleteOne deletes a single document from the specified collection based on the filter.
func (db *MongoDbClient) DeleteOne(ctx context.Context, databaseName, collectionName string, filter interface{}) (*mongo.DeleteResult, error) {
	collection := db.Mdb.Database(databaseName).Collection(collectionName)
	deleteResult, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	return deleteResult, nil
}

// DeleteMany deletes documents from the specified collection based on the filter.
func (db *MongoDbClient) DeleteMany(ctx context.Context, databaseName, collectionName string, filter interface{}) (*mongo.DeleteResult, error) {
	collection := db.Mdb.Database(databaseName).Collection(collectionName)
	deleteResult, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		return nil, err
	}
	return deleteResult, nil
}

// CountDocuments returns the number of documents that match the filter in the specified collection.
func (db *MongoDbClient) CountDocuments(ctx context.Context, databaseName, collectionName string, filter interface{}) (int64, error) {
	collection := db.Mdb.Database(databaseName).Collection(collectionName)
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Aggregate performs an aggregation query on the specified collection.
func (db *MongoDbClient) Aggregate(ctx context.Context, databaseName, collectionName string, pipeline interface{}, result interface{}) error {
	collection := db.Mdb.Database(databaseName).Collection(collectionName)
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	// Decode the result into the provided slice
	if err := cursor.All(ctx, result); err != nil {
		return err
	}

	return nil
}

// CreateIndex creates an index on the specified collection with the given keys and options.
func (db *MongoDbClient) CreateIndex(ctx context.Context, databaseName, collectionName string, keys interface{}, options *options.IndexOptions) (string, error) {
	collection := db.Mdb.Database(databaseName).Collection(collectionName)
	indexName, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    keys,
		Options: options,
	})
	if err != nil {
		return "", err
	}
	return indexName, nil
}
