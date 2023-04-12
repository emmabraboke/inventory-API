package mongoTransactionRepo

import (
	"context"
	"inventory/internals/entity/transactionEntity"
	"inventory/internals/repository/transactionRepo"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepo struct {
	transactionCollection *mongo.Collection
}

func NewTransactionRepo(transactionCollection *mongo.Collection) transactionRepo.TransactionRepository {
	return &mongoRepo{transactionCollection: transactionCollection}
}

func (t *mongoRepo) CreateTransaction(req *transactionEntity.Transaction) error {

	req.Id = primitive.NewObjectID()
	_, err := t.transactionCollection.InsertOne(context.Background(), req)

	return err
}

func (t *mongoRepo) GetTransactions() ([]*transactionEntity.Transaction, error) {
	ctx := context.TODO()

	transaction := []*transactionEntity.Transaction{}
	cursor, err := t.transactionCollection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &transaction)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err := cursor.Err(); err != nil {
		// handle any errors that occurred during iteration
		log.Println(err)
	}

	return transaction, nil
}

func (t *mongoRepo) GetTransaction(transactionId string) (*transactionEntity.Transaction, error) {
	id, _ := primitive.ObjectIDFromHex(transactionId)
	var transaction transactionEntity.Transaction
	err := t.transactionCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&transaction)

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (t *mongoRepo) UpdateTransaction(id string, req *transactionEntity.UpdateTransaction) (*transactionEntity.Transaction, error) {

	var transaction transactionEntity.Transaction
	transactionId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": transactionId}

	var update primitive.D

	if req.Status != nil {
		update = append(update, bson.E{Key: "status", Value: req.Status})
	}

	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	update = append(update, bson.E{Key: "updatedAt", Value: updatedAt})

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err := t.transactionCollection.FindOneAndUpdate(context.TODO(), filter, bson.M{"$set": update}, opts).Decode(&transaction)

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (t *mongoRepo) DeleteTransaction(id string) error {

	transactionId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": transactionId}

	_, err = t.transactionCollection.DeleteOne(context.Background(), filter)

	return err
}
