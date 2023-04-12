package mongoDatabase

import (
	"context"
	"inventory/internals/database"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type db struct {
	mongoUri     string
	databaseName string
}

func NewDatabase(mongoUri, databaseName string) database.Database{
	return &db{mongoUri: mongoUri, databaseName: databaseName}
}

func (t *db) Connect() *mongo.Client {

	clientOptions := options.Client().ApplyURI(t.mongoUri)

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	//check connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal("Database connection failed")
	}

	log.Println("Database connected successfully")

	return client

}

func (t *db) CreateCollection(client *mongo.Client, name string) *mongo.Collection {
	collection := client.Database(t.databaseName).Collection(name)

	return collection
}
