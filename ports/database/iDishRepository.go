package database

import (
	"restaurant/models"
)

//Se usar como port para acceder a la base de datos
type IDishRepository interface {
	Insert(dish models.Dish) error
	FindById(id string) (error, models.Dish)
	FindAll() (error, []models.Dish)
	DeleteById(id string) error
	UpdateById(id string, dish models.Dish) (error, models.Dish)
	//FindByDishId(id string) (error, models.Dish)

}
