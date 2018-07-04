package database

import (
	"restaurant/models"
)

type IUserRepository interface {
	Authenticate(username string, password string) bool
	Insert(user models.User) error
	FindById(id string) (error, models.User)
	FindAll() (error, []models.User)
	FindByUsername(username string) (error, models.User)
	DeleteById(id string) error
	UpdateById(id string, user models.User) (error, models.User)
}
