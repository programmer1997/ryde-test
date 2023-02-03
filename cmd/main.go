package main

import (
	"github.com/programmer1997/ryde-test/internal/db"
	"github.com/programmer1997/ryde-test/internal/router"
)

func main() {
	// Initialise DB
	db := db.NewMongoDBClient()
	// Initialise router
	router.CreateRouter(db)
}
