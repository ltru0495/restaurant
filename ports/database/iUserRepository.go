package database

import (
	"productTest/models"
)

type IUserRepository interface {
	Insert(user models.User) error
	FindById(id string) (error, models.User)
	FindAll() (error, []models.User)
	FindByUsername(username string) (error, models.User)
	DeleteById(id string) error
	UpdateById(id string, user models.User) error
}
