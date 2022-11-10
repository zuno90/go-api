package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Insance Database
var DB *mongo.Client

func ConnectDB() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
    mongoClient, err := mongo.NewClient(mongoOptions)
	if err != nil {
        log.Fatal(err)
    }
    // defer mongoClient.Disconnect(ctx)
    mongoClient.Connect(ctx)
    fmt.Println("Connected to MongoDB")
    Ping(mongoClient, ctx)

    // assign DB -> client of MONGODB
    DB = mongoClient
}
func Ping(mongoClient *mongo.Client, ctx context.Context) error {
    if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
        return err
    }
    fmt.Println("PONG, MongoDb is connect")
    return nil
}

// getting database collections
func GetCollection(mongoClient *mongo.Client, collectionName string) *mongo.Collection {
    collection := mongoClient.Database(os.Getenv("DB_NAME")).Collection(collectionName)
    return collection
}

// func CloseConnection(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
// 	defer cancel()
// 	defer func() {
// 		if err := client.Disconnect(ctx); err != nil{panic(err)}
// 	}()
// }