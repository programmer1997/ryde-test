package db

import (
	"errors"
	"github.com/programmer1997/ryde-test/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockDBClient struct {
	mockDB map[primitive.ObjectID]models.User
}

func NewMockDBClient() MockDBClient {
	return MockDBClient{mockDB: make(map[primitive.ObjectID]models.User)}
}

func (db MockDBClient) GetUserById(id string) (models.User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	var user models.User
	if err != nil {
		return user, err
	}
	if v, ok := db.mockDB[objId]; ok {
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
	if _, ok := db.mockDB[objId]; ok {
		delete(db.mockDB, objId)
		return nil
	} else {
		return errors.New("user does not exist")
	}
}
func (db MockDBClient) CreateUser(user models.User) (models.User, error) {
	objId := primitive.NewObjectID()
	user.Id = objId
	db.mockDB[objId] = user
	return user, nil
}
func (db MockDBClient) UpdateUser(id string, user models.User) (models.User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}
	if _, ok := db.mockDB[objId]; ok {
		db.mockDB[objId] = user
		return user, nil
	} else {
		return user, errors.New("user does not exist")
	}

}
