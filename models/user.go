package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Dob         time.Time          `bson:"dob,omitempty" json:"dob"`
	Address     string             `bson:"address,omitempty" json:"address"`
	Description string             `bson:"description,omitempty" json:"description"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
}
