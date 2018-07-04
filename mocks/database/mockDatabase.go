package database

import (
	"gopkg.in/mgo.v2/bson"
	"restaurant/models"
)

var MockDishDatabase []models.Dish

type MockDatabase struct {
}

func (m MockDatabase) Connect() {
	dish := models.Dish{bson.NewObjectId(), "1", "Ceviche", "ceviche", 20.0, "1.jpg"}
	MockDishDatabase = append(MockDishDatabase, dish)
	dish = models.Dish{bson.NewObjectId(), "2", "Honey Chicken", "honey chicken", 12.0, "2.jpg"}
	MockDishDatabase = append(MockDishDatabase, dish)
	dish = models.Dish{bson.NewObjectId(), "3", "Curry Extra Picante", "curry extra picante", 18.0, "3.jpg"}
	MockDishDatabase = append(MockDishDatabase, dish)
	dish = models.Dish{bson.NewObjectId(), "4", "Causa Rellena de Pollo", "causa rellena de pollo", 10.0, "7.jpg"}
	MockDishDatabase = append(MockDishDatabase, dish)
	dish = models.Dish{bson.NewObjectId(), "5", "Saltado de Pollo", "saltado de pollo", 18.0, "4.jpg"}
	MockDishDatabase = append(MockDishDatabase, dish)
	dish = models.Dish{bson.NewObjectId(), "6", "Arroz con Mariscos", "arros con mariscos", 25.0, "5.jpg"}
	MockDishDatabase = append(MockDishDatabase, dish)
}
