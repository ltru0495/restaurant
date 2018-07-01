package database

import (
	"restaurant/models"
)

type MockDishRepository struct {
	MockDatabase MockDatabase
}

func (repo MockDishRepository) InsertDish(dish models.Dish) error {
	MockDishDatabase = append(MockDishDatabase, dish)
	return nil
}

func (repo MockDishRepository) FindById(id string) (error, models.Dish) {
	dish := models.Dish{}
	return nil, dish
}

func (repo MockDishRepository) FindByDishId(did string) (error, models.Dish) {
	dish := models.Dish{}
	for _, v := range MockDishDatabase {
		if v.Dishid == did {
			return nil, v
		}
	}
	return nil, dish
}

func (repo MockDishRepository) FindAllDishes() (error, []models.Dish) {
	var dishes []models.Dish
	for _, v := range MockDishDatabase {
		dishes = append(dishes, v)
	}
	return nil, dishes
}
