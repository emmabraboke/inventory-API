package mongoUserRepo

import (
	"context"
	"inventory/internals/entity/userEntity"
	"inventory/internals/repository/userRepo"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepo struct {
	userCollection *mongo.Collection
}

func NewUserRepo(userCollection *mongo.Collection) userRepo.UserRepository {
	return &mongoRepo{userCollection: userCollection}
}

func (t *mongoRepo) CreateUser(req *userEntity.CreateUserReq) error {
	req.Id = primitive.NewObjectID()
	_, err := t.userCollection.InsertOne(context.Background(), req)

	return err
}

func (t *mongoRepo) GetUsers() ([]*userEntity.CreateUserRes, error) {
	ctx := context.TODO()

	user := []*userEntity.CreateUserRes{}
	cursor, err := t.userCollection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &user)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err := cursor.Err(); err != nil {
		// handle any errors that occurred during iteration
		log.Println(err)
	}

	return user, nil
}

func (t *mongoRepo) GetUser(userId string) (*userEntity.CreateUserRes, error) {
	id, _ := primitive.ObjectIDFromHex(userId)
	var user userEntity.CreateUserRes
	err := t.userCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (t *mongoRepo) GetUserByEmail(email string) (*userEntity.CreateUserReq, error) {
	var user userEntity.CreateUserReq
	err := t.userCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (t *mongoRepo) UpdateUser(id string, req *userEntity.UpdateUserReq) (*userEntity.CreateUserRes, error) {

	var user userEntity.CreateUserRes
	userId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": userId}

	var update primitive.D

	if req.FirstName != nil {
		update = append(update, bson.E{Key: "firstName", Value: req.FirstName})
	}

	if req.LastName != nil {
		update = append(update, bson.E{Key: "lastName", Value: req.LastName})
	}

	if req.ProfileImage != nil {
		update = append(update, bson.E{Key: "profileImage", Value: req.ProfileImage})
	}

	if req.RefreshToken != nil {
		update = append(update, bson.E{Key: "refreshToken", Value: req.RefreshToken})
	}

	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	update = append(update, bson.E{Key: "updatedAt", Value: updatedAt})

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err := t.userCollection.FindOneAndUpdate(context.TODO(), filter, bson.M{"$set": update}, opts).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (t *mongoRepo) DeleteUser(id string) error {

	userId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": userId}

	_, err = t.userCollection.DeleteOne(context.Background(), filter)

	return err
}
