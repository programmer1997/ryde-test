package router

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/programmer1997/ryde-test/internal/db"
	"github.com/programmer1997/ryde-test/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserSuccess(t *testing.T) {
	id := primitive.NewObjectID()
	db := db.NewMockDBClient(
		map[primitive.ObjectID]models.User{
			id: {Id: id, Name: "AJ", Address: "1 street", Description: ""},
		})
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: id.Hex()}}
	getUserById(c, db)
	expectedBody := fmt.Sprintf("{\"id\":\"%s\",\"name\":\"AJ\",\"dob\":\"0001-01-01T00:00:00Z\",\"address\":\"1 street\",\"description\":\"\",\"createdAt\":\"0001-01-01T00:00:00Z\"}", id.Hex())
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedBody, w.Body.String())
}

func TestGetUserNotExist(t *testing.T) {
	id := primitive.NewObjectID()
	notExistId := primitive.NewObjectID()
	db := db.NewMockDBClient(
		map[primitive.ObjectID]models.User{
			id: {Id: id, Name: "AJ", Address: "1 street", Description: ""},
		})
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: notExistId.Hex()}}
	getUserById(c, db)
	expectedBody := "{\"error\":\"user does not exist\"}"
	assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	assert.Equal(t, expectedBody, w.Body.String())
}

func TestCreateUserSuccess(t *testing.T) {
	db := db.NewMockDBClient(make(map[primitive.ObjectID]models.User))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("{\"name\":\"AJ\",\"address\":\"31 street\"}")))
	createUser(c, db)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateInvalidRequest(t *testing.T) {
	db := db.NewMockDBClient(make(map[primitive.ObjectID]models.User))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodDelete, "", bytes.NewBuffer([]byte("{{}")))
	deleteUser(c, db)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateUserSuccess(t *testing.T) {
	id := primitive.NewObjectID()
	db := db.NewMockDBClient(
		map[primitive.ObjectID]models.User{
			id: {Id: id, Name: "AJ", Address: "1 street", Description: ""},
		})
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: id.Hex()}}
	c.Request, _ = http.NewRequest(http.MethodPut, "", bytes.NewBuffer([]byte(fmt.Sprintf("{\"id\":\"%s\",\"name\":\"AABGB\",\"dob\":\"0001-01-01T00:00:00Z\",\"address\":\"2 street\",\"description\":\"\",\"createdAt\":\"0001-01-01T00:00:00Z\"}", id.Hex()))))
	updateUser(c, db)
	expectedBody := fmt.Sprintf("{\"id\":\"%s\",\"name\":\"AABGB\",\"dob\":\"0001-01-01T00:00:00Z\",\"address\":\"2 street\",\"description\":\"\",\"createdAt\":\"0001-01-01T00:00:00Z\"}", id.Hex())
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedBody, w.Body.String())
}

func TestUpdateUserNotExist(t *testing.T) {
	id := primitive.NewObjectID()
	idNotExist := primitive.NewObjectID()
	db := db.NewMockDBClient(
		map[primitive.ObjectID]models.User{
			id: {Id: id, Name: "AJ", Address: "1 street", Description: ""},
		})
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: idNotExist.Hex()}}
	c.Request, _ = http.NewRequest(http.MethodPut, "", bytes.NewBuffer([]byte(fmt.Sprintf("{\"id\":\"%s\",\"name\":\"AABGB\",\"dob\":\"0001-01-01T00:00:00Z\",\"address\":\"2 street\",\"description\":\"\",\"createdAt\":\"0001-01-01T00:00:00Z\"}", id.Hex()))))
	updateUser(c, db)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestDeleteUser(t *testing.T) {
	id := primitive.NewObjectID()
	db := db.NewMockDBClient(
		map[primitive.ObjectID]models.User{
			id: {Id: id, Name: "AJ", Address: "1 street", Description: ""},
		})
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: id.Hex()}}
	deleteUser(c, db)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, len(db.MockDb), 0)
}

func TestDeleteUserDoesNotExist(t *testing.T) {
	id := primitive.NewObjectID()
	db := db.NewMockDBClient(
		map[primitive.ObjectID]models.User{})
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: id.Hex()}}
	deleteUser(c, db)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
