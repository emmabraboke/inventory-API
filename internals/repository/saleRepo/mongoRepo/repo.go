package mongoSaleRepo

import (
	"context"
	"inventory/internals/entity/saleEntity"
	"inventory/internals/repository/saleRepo"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepo struct {
	saleCollection *mongo.Collection
}

func NewSaleRepo(saleCollection *mongo.Collection) saleRepo.SaleRepository {
	return &mongoRepo{saleCollection: saleCollection}
}

func (t *mongoRepo) CreateSale(req *saleEntity.Sale) error {

	req.Id = primitive.NewObjectID()

	_, err := t.saleCollection.InsertOne(context.Background(), req)

	return err
}

func (t *mongoRepo) GetSales() ([]*saleEntity.Sale, error) {
	ctx := context.TODO()

	sale := []*saleEntity.Sale{}
	cursor, err := t.saleCollection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &sale)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err := cursor.Err(); err != nil {
		// handle any errors that occurred during iteration
		log.Println(err)
	}

	return sale, nil
}

func (t *mongoRepo) GetSale(saleId string) (*saleEntity.Sale, error) {
	id, _ := primitive.ObjectIDFromHex(saleId)
	var sale saleEntity.Sale
	err := t.saleCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&sale)

	if err != nil {
		return nil, err
	}

	return &sale, nil
}

func (t *mongoRepo) DeleteSale(id string) error {

	saleId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": saleId}

	_, err = t.saleCollection.DeleteOne(context.Background(), filter)

	return err
}
