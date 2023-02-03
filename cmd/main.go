package main

import (
	"github.com/gin-gonic/gin"
	"github.com/programmer1997/ryde-test/internal/db"
	router2 "github.com/programmer1997/ryde-test/internal/router"
	"log"
)

var router *gin.Engine

func init() {
	db := db.NewMongoDBClient()

	// Initialise router
	router = router2.CreateRouter(db)
}

func main() {
	log.Println("Starting server")
}
