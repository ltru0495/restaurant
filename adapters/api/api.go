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

	UserRepository dbPorts.IUserRepository

	MenuRepository dbPorts.IMenuRepository
}

func (api API) PostDishHandler(w http.ResponseWriter, r *http.Request) {
	var dish models.Dish
	err := json.NewDecoder(r.Body).Decode(&dish)
	if err != nil {
		log.Println(err)
		return
	}
	dish.ID = bson.NewObjectId()
	err = api.DishRepository.Insert(dish)
	if err != nil {
		log.Println(err)
		return
	}
	//json.NewEncoder(w).Encode(dish)
	j, err := json.Marshal(dish)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
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
		log.Println(err)
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
		log.Println(err)
	}

	w.WriteHeader(http.StatusNoContent)
}

func (api API) PutDishHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var dish models.Dish
	err := json.NewDecoder(r.Body).Decode(&dish)
	if err != nil {
		log.Println(err)
	}

	err, updatedDish := api.DishRepository.UpdateById(id, dish)
	if err != nil {
		log.Println(err)
	}

	j, err := json.Marshal(updatedDish)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if (models.Dish{} == updatedDish) {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(j)
}

func (api API) PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		return
	}
	user.ID = bson.NewObjectId()
	err = api.UserRepository.Insert(user)
	if err != nil {
		log.Println(err)
		return
	}
	j, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)

}

func (api API) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	err, users := api.UserRepository.FindAll()
	if err != nil {
		return
	}
	//json.NewEncoder(w).Encode(dishes)
	j, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func (api API) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	err, user := api.UserRepository.FindById(id)
	//No encontrado
	if err != nil {
		user = models.User{}
	}

	j, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if (models.User{} == user) {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Write(j)
}

func (api API) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	err := api.UserRepository.DeleteById(id)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusNoContent)
}

func (api API) PutUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}

	err, updatedUser := api.UserRepository.UpdateById(id, user)
	if err != nil {
		log.Println(err)
	}

	j, err := json.Marshal(updatedUser)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if (models.User{} == updatedUser) {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(j)
}

func (api API) PostMenuHandler(w http.ResponseWriter, r *http.Request) {
	var menu models.Menu
	err := json.NewDecoder(r.Body).Decode(&menu)
	if err != nil {
		log.Println(err)
		return
	}
	menu.ID = bson.NewObjectId()
	err = api.MenuRepository.Insert(menu)

	j, err := json.Marshal(menu)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)

}

func (api API) GetAllMenuHandler(w http.ResponseWriter, r *http.Request) {
	var menus []models.Menu

	err, menus := api.MenuRepository.FindAll()
	if err != nil {
		log.Println(err)
		return
	}

	j, err := json.Marshal(menus)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func (api API) GetDailyMenuHandler(w http.ResponseWriter, r *http.Request) {
	var dishes []models.Dish
	var menu models.Menu

	params := mux.Vars(r)
	date := params["date"]

	err, menu := api.MenuRepository.FindByDate(date)
	if err != nil {
		log.Println(err)
		return
	}

	var aux models.Dish
	for _, v := range menu.Dishes {
		err, aux = api.DishRepository.FindById(v)
		dishes = append(dishes, aux)
		println(aux.String())
	}
}
