package db

import (
	"context"
	"fmt"
	"log"

	"github.com/sirawong/bookrental-microservices/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB(ctx context.Context, config config.DatabaseConfig) *mongo.Client {
	fmt.Println("Connecting to MongoDB...")

	dbURI := fmt.Sprintf("mongodb://%s:%s@mongodb:%s/?authSource=admin&ssl=false", config.DBUsername, config.DBPassword, config.DBPort)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	} else {
		fmt.Println("Connected to MongoDB")
	}

	// go func() {
	// 	client.Disconnect(context.Background())
	// 	fmt.Println("MongoDB connection closed")
	// }()

	return client
}
