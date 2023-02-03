package router

import (
	"github.com/gin-gonic/gin"
	"github.com/programmer1997/ryde-test/internal/db"
	"github.com/programmer1997/ryde-test/models"
	"net/http"
)

const (
	idMissingError   = "Request URL should contain Id"
	successDeleteMsg = "Successfully Deleted"
)

func CreateRouter(client db.DBClient) {
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
	// Delete user
	r.DELETE("/v1/users/delete/:id", func(c *gin.Context) {
		deleteUser(c, client)
	})
	r.Run()
}

func getUserById(c *gin.Context, client db.DBClient) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": idMissingError})
		return
	}
	res, err := client.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func createUser(c *gin.Context, client db.DBClient) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := client.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

func updateUser(c *gin.Context, client db.DBClient) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": idMissingError})
		return
	}
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res, err := client.UpdateUser(id, user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func deleteUser(c *gin.Context, client db.DBClient) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": idMissingError})
		return
	}
	err := client.DeleteUser(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": successDeleteMsg})
}
