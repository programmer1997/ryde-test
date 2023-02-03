package db

import "github.com/programmer1997/ryde-test/models"

type DBClient interface {
	GetUserById(id string) (models.User, error)
	DeleteUser(id string) error
	CreateUser(user models.User) (models.User, error)
	UpdateUser(id string, user models.User) (models.User, error)
}
