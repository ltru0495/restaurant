package database

import (
	"errors"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	models "restaurant/models"
	"strconv"
)

type DishRepository struct {
	DB *mgo.Database
}

const (
	DISHCOLLECTION = "dishes"
)

func (repo DishRepository) Insert(dish models.Dish) error {
	n, err := repo.DB.C(DISHCOLLECTION).Count()
	n = n + 1
	if err != nil {
		dish.Dishid = ""
	}
	dish.Dishid = strconv.Itoa(n)
	err = repo.DB.C(DISHCOLLECTION).Insert(&dish)
	return err

}

func (repo DishRepository) FindById(did string) (error, models.Dish) {
	dish := models.Dish{}
	err := repo.DB.C(DISHCOLLECTION).Find(bson.M{"dishid": did}).One(&dish)
	if (models.Dish{} == dish) {
		return errors.New("No Encontrado"), dish
	}
	return err, dish
}

func (repo DishRepository) FindAll() (error, []models.Dish) {
	var dishes []models.Dish
	aux := &models.Dish{}

	iter := repo.DB.C(DISHCOLLECTION).Find(nil).Iter()
	for iter.Next(&aux) {
		dishes = append(dishes, *aux)
	}
	return nil, dishes
}

func (repo DishRepository) DeleteById(id string) error {
	err := repo.DB.C(DISHCOLLECTION).Remove(bson.M{"dishid": id})
	return err
}

func (repo DishRepository) UpdateById(id string, dish models.Dish) (error, models.Dish) {
	oldDish := models.Dish{}
	err := repo.DB.C(DISHCOLLECTION).Find(bson.M{"dishid": id}).One(&oldDish)
	if (models.Dish{} == oldDish) {
		return errors.New("No Encontrado"), dish
	}

	dish.ID = oldDish.ID

	if dish.Dishid == "" {
		dish.Dishid = oldDish.Dishid
	}
	if dish.Name == "" {
		dish.Name = oldDish.Name
	}
	if dish.Description == "" {
		dish.Description = oldDish.Description
	}
	if dish.Price == 0.0 {
		dish.Price = oldDish.Price
	}
	err = repo.DB.C(DISHCOLLECTION).Update(bson.M{"dishid": id}, &dish)
	return err, dish
}

/*func (repo DishRepository) FindById(id string) (error, models.Dish) {
	dish := models.Dish{}

	// Se tiene un id valido para busqueda
	if bson.IsObjectIdHex(id) {

		// TODO
		// agregar custom error de producto no encontrado
		err := repo.DB.C(DISHCOLLECTION).FindId(bson.ObjectIdHex(id)).One(&dish)
		return err, dish
	} else {
		// TODO
		println("TODO ID NO VALIDO")
		retString := "ObjectId No Valido"
		return errors.New(retString), dish
	}
}*/
