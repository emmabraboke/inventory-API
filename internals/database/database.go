package database

import "go.mongodb.org/mongo-driver/mongo"

type Database interface {
	Connect() *mongo.Client
	CreateCollection(client *mongo.Client, name string) *mongo.Collection 
}