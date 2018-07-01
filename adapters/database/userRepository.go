package database

import (
	"errors"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	models "restaurant/models"
)

type UserRepository struct {
	DB *mgo.Database
}

const (
	USERCOLLECTION = "users"
)

func (repo UserRepository) Insert(user models.User) error {
	err := repo.DB.C(USERCOLLECTION).Insert(&user)
	return err
}

func (repo UserRepository) FindById(uid string) (error, models.User) {
	user := models.User{}
	err := repo.DB.C(USERCOLLECTION).Find(bson.M{"userid": uid}).One(&user)
	if (models.User{} == user) {
		return errors.New("No Encontrado"), user
	}
	return err, user
}

func (repo UserRepository) FindAll() (error, []models.User) {
	var users []models.User
	aux := &models.User{}

	iter := repo.DB.C(USERCOLLECTION).Find(nil).Iter()
	for iter.Next(&aux) {
		users = append(users, *aux)
	}
	return nil, users
}

func (repo UserRepository) FindByUsername(username string) (error, models.User) {
	user := models.User{}
	err := repo.DB.C(USERCOLLECTION).Find(bson.M{"username": username}).One(&user)
	if (models.User{} == user) {
		return errors.New("No Encontrado"), user
	}
	return err, user
}

func (repo UserRepository) DeleteById(id string) error {
	err := repo.DB.C(USERCOLLECTION).Remove(bson.M{"userid": id})
	return err
}

func (repo UserRepository) UpdateById(id string, user models.User) (error, models.User) {
	return nil, models.User{}
}
