package db

import (
	"errors"
	"github.com/programmer1997/ryde-test/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MockDBClient stores user information in-memory and is used for unit testing
type MockDBClient struct {
	MockDb map[primitive.ObjectID]models.User
}

func NewMockDBClient(db map[primitive.ObjectID]models.User) MockDBClient {
	return MockDBClient{MockDb: db}
}

func (db MockDBClient) GetUserById(id string) (models.User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	var user models.User
	if err != nil {
		return user, err
	}
	if v, ok := db.MockDb[objId]; ok {
		return v, nil
	} else {
		return user, errors.New("user does not exist")
	}
}
func (db MockDBClient) DeleteUser(id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	if _, ok := db.MockDb[objId]; ok {
		delete(db.MockDb, objId)
		return nil
	} else {
		return errors.New("user does not exist")
	}
}
func (db MockDBClient) CreateUser(user models.User) (models.User, error) {
	objId := primitive.NewObjectID()
	user.Id = objId
	db.MockDb[objId] = user
	return user, nil
}
func (db MockDBClient) UpdateUser(id string, user models.User) (models.User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}
	if _, ok := db.MockDb[objId]; ok {
		db.MockDb[objId] = user
		return user, nil
	} else {
		return user, errors.New("user does not exist")
	}

}
