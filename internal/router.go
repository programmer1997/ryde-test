package internal

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/programmer1997/ryde-test/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

func CreateRouter(client *mongo.Client) *gin.Engine {
	r := gin.Default()

	// Get user by id
	r.GET("/v1/users/:id", func(c *gin.Context) {
		getUserById(c, client)
	})
	// Create user
	r.POST("/v1/users/create", func(c *gin.Context) {
		createUser(c, client)

	})
	// Update user
	r.PUT("/v1/users/update/:id", func(c *gin.Context) {
		updateUser(c, client)
	})
	r.DELETE("/v1/users/delete/:id", func(c *gin.Context) {
		deleteUser(c, client)
	})
	r.Run()
	return r
}

func getUserById(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")
	if id == "" {
		fmt.Print("id not found")

	}
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Print("error while converting to object id  ")

	}
	var user models.User
	user.Id = objId
	collection := client.Database("rydedb").Collection("users")
	collection.Find(context.TODO(), bson.M{})
	err = collection.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&user)
	if err != nil {
		fmt.Print("error while fetching from db ")
	}
	fmt.Print(user)
	c.JSON(http.StatusOK, user)
}

func createUser(c *gin.Context, client *mongo.Client) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Id = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	collection := client.Database("rydedb").Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func updateUser(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	objId, err := primitive.ObjectIDFromHex(id)
	user.Id = objId
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
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
	collection := client.Database("rydedb").Collection("users")
	result, err := collection.UpdateOne(nil, filter, update)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if result.ModifiedCount == 0 {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}

func deleteUser(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	collection := client.Database("rydedb").Collection("users")
	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objId})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if result.DeletedCount == 0 {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, "Successfully deleted")

}
