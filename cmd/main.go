package main

import (
	"github.com/gin-gonic/gin"
	"github.com/programmer1997/ryde-test/internal"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var client *mongo.Client
var router *gin.Engine

func init() {
	// Initialise DB
	c, err := internal.InitDB()
	if err != nil {
		log.Fatal("DB Initialisation failed. Shutting down")
	}
	client = c

	// Initialise router
	router = internal.CreateRouter(client)
}

func main() {
	log.Println("Starting server")
}
