package api

import (
	"encoding/json"
	// "fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"restaurant/models"
	dbPorts "restaurant/ports/database"
)

type API struct {
	DishRepository dbPorts.IDishRepository
}

func (api API) PostDishHandler(w http.ResponseWriter, r *http.Request) {
	var dish models.Dish
	err := json.NewDecoder(r.Body).Decode(&dish)
	if err != nil {
		log.Fatal(err)
		return
	}
	dish.ID = bson.NewObjectId()
	err = api.DishRepository.Insert(dish)
	if err != nil {
		log.Fatal(err)
		return
	}
	//json.NewEncoder(w).Encode(dish)
	j, err := json.Marshal(dish)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)

}

func (api API) GetAllDishesHandler(w http.ResponseWriter, r *http.Request) {
	err, dishes := api.DishRepository.FindAll()
	if err != nil {
		return
	}
	//json.NewEncoder(w).Encode(dishes)
	j, err := json.Marshal(dishes)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func (api API) GetDishHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	err, dish := api.DishRepository.FindById(id)
	//No encontrado
	if err != nil {
		dish = models.Dish{}
	}

	j, err := json.Marshal(dish)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if (models.Dish{} == dish) {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Write(j)
}

func (api API) DeleteDishHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	err := api.DishRepository.DeleteById(id)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusNoContent)
}

func (api API) PutDishHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var dish models.Dish
	err := json.NewDecoder(r.Body).Decode(&dish)
	if err != nil {
		log.Fatal(err)
	}

	err, updatedDish := api.DishRepository.UpdateById(id, dish)
	if err != nil {
		log.Fatal(err)
	}

	j, err := json.Marshal(updatedDish)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if (models.Dish{} == updatedDish) {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(j)
}
