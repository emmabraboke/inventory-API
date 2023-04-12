package mongoProductRepo

import (
	"context"
	"inventory/internals/entity/productEntity"
	"inventory/internals/repository/productRepo"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepo struct {
	productCollection *mongo.Collection
}

func NewproductRepo(productCollection *mongo.Collection) productRepo.ProductRepository {
	return &mongoRepo{productCollection: productCollection}
}

func (t *mongoRepo) CreateProduct(req *productEntity.Product) error {

	req.Id = primitive.NewObjectID()

	_, err := t.productCollection.InsertOne(context.Background(), req)

	return err
}

func (t *mongoRepo) GetProducts() ([]*productEntity.Product, error) {
	ctx := context.TODO()

	product := []*productEntity.Product{}
	cursor, err := t.productCollection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &product)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err := cursor.Err(); err != nil {
		// handle any errors that occurred during iteration
		log.Println(err)
	}

	return product, nil
}

func (t *mongoRepo) GetProduct(productId string) (*productEntity.Product, error) {
	id, _ := primitive.ObjectIDFromHex(productId)
	var product productEntity.Product
	err := t.productCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&product)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (t *mongoRepo) UpdateProduct(id string, req *productEntity.UpdateProduct) (*productEntity.Product, error) {

	var product productEntity.Product
	productId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": productId}

	var update primitive.D

	if req.Name != nil {
		update = append(update, bson.E{Key: "name", Value: req.Name})
	}

	if req.Price != nil {
		update = append(update, bson.E{Key: "price", Value: req.Price})
	}

	if req.Quantity != nil {
		update = append(update, bson.E{Key: "quantity", Value: req.Quantity})
	}

	if req.Description != nil {
		update = append(update, bson.E{Key: "description", Value: req.Description})
	}

	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	update = append(update, bson.E{Key: "updatedAt", Value: updatedAt})

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err := t.productCollection.FindOneAndUpdate(context.TODO(), filter, bson.M{"$set": update}, opts).Decode(&product)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (t *mongoRepo) UpdateProductQuantity(id string, req *productEntity.UpdateProduct) (*productEntity.Product, error) {

	var product productEntity.Product
	productId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": productId}

	var update primitive.D

	if req.Quantity != nil {
		update = append(update, bson.E{Key: "quantity", Value: req.Quantity})
	}

	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	update = append(update, bson.E{Key: "updatedAt", Value: updatedAt})

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err := t.productCollection.FindOneAndUpdate(context.TODO(), filter, bson.M{"$inc": update}, opts).Decode(&product)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (t *mongoRepo) DeleteProduct(id string) error {

	productId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": productId}

	_, err = t.productCollection.DeleteOne(context.Background(), filter)

	return err
}
