package main

import (
	"fmt"
	apiAdapter "restaurant/adapters/api"
	dbAdapters "restaurant/adapters/database"
	mocksDb "restaurant/mocks/database"
	apiPort "restaurant/ports/api"
	dbPorts "restaurant/ports/database"
	//"strconv"
	//"time"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	daoDb          dbPorts.IDatabase
	dishRepository dbPorts.IDishRepository
	userRepository dbPorts.IUserRepository
	api            apiPort.IAPI

	/*MOCKS*/
	mockDatabase       mocksDb.MockDatabase
	mockDishRepository mocksDb.MockDishRepository
)

func init() {

	/* 	Conexion a la base datos*/
	url := "127.0.0.1:27017"
	db := "restaurant"
	daoDb = &dbAdapters.MongoDbAdapter{url, db}
	daoDb.Connect()

	// Se tiene una variable compartida Dao en el paquete adapters/database
	// esta variable se pasara a los repositorios para que sean almacenados como miembros
	dishRepository = &dbAdapters.DishRepository{dbAdapters.Dao}
	userRepository = &dbAdapters.UserRepository{dbAdapters.Dao}

	// 	API //
	api = apiAdapter.API{dishRepository}

	// MOCKS //
	/*mockDatabase = mocksDb.MockDatabase{}
	mockDatabase.Connect()
	mockDishRepository = mocksDb.MockDishRepository{mockDatabase}

	api = apiAdapter.Api{mockDishRepository}*/
}

func main() {
	//var productRepo pPort.ProductRepository

	/*var buy models.Buy
	ps := []string{"2", "1", "3"}
	buy.UserId = "lt1235"
	buy.Dishes = ps
	buy.DoneAt = time.Now()
	buy.Total = 12.3

	fmt.Println(strconv.FormatFloat(a, 'f', 4, 64))
	fmt.Println("************COMPRA**********")
	for k, v := range buy.Dishes {
		fmt.Println("Producto " + strconv.Itoa(k))
		err, prodAux := dishRepository.FindByDishId(v)
		fmt.Println(prodAux)
		if err != nil {
			fmt.Println(err)
		}
	}
	*/
	fmt.Println("*********")
	r := mux.NewRouter()
	r.HandleFunc("/api/dishes", api.GetAllDishesHandler).Methods("GET")
	r.HandleFunc("/api/dishes/{id:[0-9]+}", api.GetDishHandler).Methods("GET")
	r.HandleFunc("/api/dishes", api.PostDishHandler).Methods("POST")
	r.HandleFunc("/api/dishes/{id:[0-9]+}", api.DeleteDishHandler).Methods("DELETE")
	r.HandleFunc("/api/dishes/{id:[0-9]+}", api.PutDishHandler).Methods("PUT")
	err := http.ListenAndServe(":9090", r)
	if err != nil {
		log.Fatal(err)
	}
}
