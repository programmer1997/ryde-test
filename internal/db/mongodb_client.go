package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/programmer1997/ryde-test/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// TODO: Handle config separately using Cobra (https://github.com/spf13/cobra)
const (
	MongoDBURI     = "mongodb://localhost:27017"
	DbName         = "rydedb"
	CollectionName = "users"
)

func InitDB() (client *mongo.Client, err error) {
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoDBURI))
	return
}

type MongoDBClient struct {
	client *mongo.Client
}

func NewMongoDBClient() MongoDBClient {
	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoDBURI))
	if err != nil {
		log.Fatal("Could not connect to Mongo DB")
	}
	return MongoDBClient{client: c}
}

func (db MongoDBClient) GetUserById(id string) (models.User, error) {
	collection := db.client.Database(DbName).Collection(CollectionName)
	var user models.User
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}
	res := collection.FindOne(context.TODO(), bson.M{"_id": objId})
	fmt.Print(res.Err())
	if res.Err() != nil {
		return user, res.Err()
	} else {
		res.Decode(&user)
	}
	return user, err
}

func (db MongoDBClient) CreateUser(user models.User) (models.User, error) {
	collection := db.client.Database(DbName).Collection(CollectionName)
	user.Id = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	_, err := collection.InsertOne(context.TODO(), user)
	return user, err
}

func (db MongoDBClient) UpdateUser(id string, user models.User) (models.User, error) {
	collection := db.client.Database(DbName).Collection(CollectionName)
	objId, err := primitive.ObjectIDFromHex(id)
	user.Id = objId
	if err != nil {
		return user, err
	}
	filter := bson.M{"_id": objId}
	update := bson.M{
		"$set": bson.M{
			"id":          user.Id,
			"name":        user.Name,
			"dob":         user.Dob,
			"address":     user.Address,
			"description": user.Description,
		},
	}
	result, err := collection.UpdateOne(nil, filter, update)
	if err != nil {
		return user, err
	}
	if result.ModifiedCount == 0 {
		err = errors.New("userId does not exist")
	}

	return user, err
}

func (db MongoDBClient) DeleteUser(id string) error {
	collection := db.client.Database(DbName).Collection(CollectionName)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objId})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("user id does not exist")
	}
	return nil
}
