package database

import (
	"errors"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"restaurant/models"
)

type MenuRepository struct {
	DB *mgo.Database
}

const (
	MENUCOLLECTION = "menu"
)

func (repo MenuRepository) Insert(menu models.Menu) error {
	err := repo.DB.C(MENUCOLLECTION).Insert(&menu)
	return err
}

func (repo MenuRepository) FindAll() (error, []models.Menu) {
	var menu []models.Menu
	aux := &models.Menu{}

	iter := repo.DB.C(MENUCOLLECTION).Find(nil).Iter()
	for iter.Next(&aux) {
		menu = append(menu, *aux)
	}
	return nil, menu
}

func (repo MenuRepository) FindByDate(date string) (error, models.Menu) {
	var menu models.Menu
	err := repo.DB.C(MENUCOLLECTION).Find(bson.M{"date": date}).One(&menu)
	if menu.Date == "" {
		return errors.New("No encontrado"), menu
	}
	return err, menu
}
