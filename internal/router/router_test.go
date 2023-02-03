package router

import (
	"github.com/gin-gonic/gin"
	"github.com/programmer1997/ryde-test/internal/db"
	"github.com/programmer1997/ryde-test/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetUser(t *testing.T) {
	id := primitive.NewObjectID()
	db := db.NewMockDBClient(
		map[primitive.ObjectID]models.User{
			id: {Id: id, Name: "AJ", Address: "1 street", Dob: time.Now(), Description: ""},
		})
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{gin.Param{Key: "id", Value: id.Hex()}}
	getUserById(c, db)

}

func TestCreateUser(t *testing.T) {

}

func TestUpdateUser(t *testing.T) {

}

func TestDeleteUser(t *testing.T) {

}
