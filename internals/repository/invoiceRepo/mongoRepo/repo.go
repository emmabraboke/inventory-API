package mongoinvoiceRepo

import (
	"context"
	"inventory/internals/entity/invoiceEntity"
	"inventory/internals/repository/invoiceRepo"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepo struct {
	invoiceCollection *mongo.Collection
}

func NewInvoiceRepo(invoiceCollection *mongo.Collection) invoiceRepo.InvoiceRepository {
	return &mongoRepo{invoiceCollection: invoiceCollection}
}

func (t *mongoRepo) CreateInvoice(req *invoiceEntity.Invoice) (*primitive.ObjectID, error) {

	req.Id = primitive.NewObjectID()
	invoice, err := t.invoiceCollection.InsertOne(context.Background(), req)

	if err != nil {
		return nil, err
	}

	id := invoice.InsertedID.(primitive.ObjectID)

	return &id, nil
}

func (t *mongoRepo) GetInvoices() ([]*invoiceEntity.Invoice, error) {
	ctx := context.TODO()

	invoice := []*invoiceEntity.Invoice{}
	cursor, err := t.invoiceCollection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &invoice)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err := cursor.Err(); err != nil {
		// handle any errors that occurred during iteration
		log.Println(err)
	}

	return invoice, nil
}

func (t *mongoRepo) GetInvoice(invoiceId string) (*invoiceEntity.Invoice, error) {
	id, _ := primitive.ObjectIDFromHex(invoiceId)
	var invoice invoiceEntity.Invoice
	err := t.invoiceCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&invoice)

	if err != nil {
		return nil, err
	}

	return &invoice, nil
}

func (t *mongoRepo) UpdateInvoice(id string, req *invoiceEntity.UpdateInvoice) (*invoiceEntity.Invoice, error) {

	var invoice invoiceEntity.Invoice
	invoiceId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": invoiceId}

	var update primitive.D

	if req.IsPaid != nil {
		update = append(update, bson.E{Key: "firstName", Value: req.IsPaid})
	}

	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	update = append(update, bson.E{Key: "updatedAt", Value: updatedAt})

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err := t.invoiceCollection.FindOneAndUpdate(context.TODO(), filter, bson.M{"$set": update}, opts).Decode(&invoice)

	if err != nil {
		return nil, err
	}

	return &invoice, nil
}

func (t *mongoRepo) DeleteInvoice(id string) error {

	invoiceId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": invoiceId}

	_, err = t.invoiceCollection.DeleteOne(context.Background(), filter)

	return err
}
