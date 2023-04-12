package mongoCustomerRepo

import (
	"context"
	"inventory/internals/entity/customerEntity"
	"inventory/internals/repository/customerRepo"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepo struct {
	customerCollection *mongo.Collection
}

func NewCustomerRepo(customerCollection *mongo.Collection) customerRepo.CustomerRepository {
	return &mongoRepo{customerCollection: customerCollection}
}

func (t *mongoRepo) CreateCustomer(req *customerEntity.Customer) error {
	req.Id = primitive.NewObjectID()
	_, err := t.customerCollection.InsertOne(context.Background(), req)

	return err
}

func (t *mongoRepo) GetCustomers() ([]*customerEntity.Customer, error) {
	ctx := context.TODO()

	customer := []*customerEntity.Customer{}
	cursor, err := t.customerCollection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &customer)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err := cursor.Err(); err != nil {
		// handle any errors that occurred during iteration
		log.Println(err)
	}

	return customer, nil
}

func (t *mongoRepo) GetCustomer(customerId string) (*customerEntity.Customer, error) {
	id, _ := primitive.ObjectIDFromHex(customerId)
	var customer customerEntity.Customer
	err := t.customerCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&customer)

	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (t *mongoRepo) GetCustomerByEmail(email string) (*customerEntity.Customer, error) {
	var customer customerEntity.Customer
	err := t.customerCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&customer)

	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (t *mongoRepo) UpdateCustomer(id string, req *customerEntity.UpdateCustomer) (*customerEntity.Customer, error) {

	var customer customerEntity.Customer
	customerId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": customerId}

	var update primitive.D

	if req.FirstName != nil {
		update = append(update, bson.E{Key: "firstName", Value: req.FirstName})
	}

	if req.LastName != nil {
		update = append(update, bson.E{Key: "lastName", Value: req.LastName})
	}

	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	update = append(update, bson.E{Key: "updatedAt", Value: updatedAt})

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err := t.customerCollection.FindOneAndUpdate(context.TODO(), filter, bson.M{"$set": update}, opts).Decode(&customer)

	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (t *mongoRepo) DeleteCustomer(id string) error {

	customerId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": customerId}

	_, err = t.customerCollection.DeleteOne(context.Background(), filter)

	return err
}
