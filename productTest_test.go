package main

import (
	"gopkg.in/mgo.v2/bson"
	"productTest/models"
	"testing"
	"time"
)

type MockProductsRepository struct {
}

func (repo MockProductsRepository) FindById(id string) (error, models.Product) {
	return nil, models.Product{}
}

func (repo MockProductsRepository) FindByProductId(pid string) (error, models.Product) {

	var products [6]models.Product
	products[0] = models.Product{bson.NewObjectId(), "1", "Ceviche", "ceviche", 20.0}
	products[1] = models.Product{bson.NewObjectId(), "2", "Honey Chicken", "honey chicken", 12.0}
	products[2] = models.Product{bson.NewObjectId(), "3", "Curry Extra Picante", "curry extra picante", 18.0}
	products[3] = models.Product{bson.NewObjectId(), "4", "Causa Rellena de Pollo", "causa rellena de pollo", 10.0}
	products[4] = models.Product{bson.NewObjectId(), "5", "Saltado de Pollo", "saltado de pollo", 18.0}
	products[5] = models.Product{bson.NewObjectId(), "6", "Arroz con Mariscos", "arros con mariscos", 25.0}

	for k, v := range products {
		if v.Productid == pid {
			return nil, products[k]
		}
	}
	return nil, models.Product{}
}

func TestCalculateCost(t *testing.T) {
	var mockRepo MockProductsRepository

	var buys [2]models.Buy
	expected := []float64{50.0, 0.1}

	ps := []string{"2", "1", "3"}
	buys[0].UserId = "lt1235"
	buys[0].Products = ps
	buys[0].DoneAt = time.Now()

	buys[1].Products = []string{}
	buys[1].UserId = "lt1235"
	buys[1].DoneAt = time.Now()

	for k, _ := range buys {
		t.Run("", func(t *testing.T) {
			cost := CalculateCost(buys[k], mockRepo)
			if cost != expected[k] {
				t.Errorf("Valor esperado: %f\nValor obtenido: %f", expected[k], cost)
				t.FailNow()
			}
		})

	}

}

/*{ "_id" : ObjectId("5b144279d7e0df4308b72874"), "productid" : "cev01", "name" : "Ceviche", "price" : 20, "description" : "ceviche" }
{ "_id" : ObjectId("5b144279d7e0df4308b72875"), "productid" : "pna05", "name" : "Pollo a la Naranja", "price" : 15, "description" : "pollo a la naranja" }
{ "_id" : ObjectId("5b144279d7e0df4308b72876"), "productid" : "crp04", "name" : "Causa Rellena de Pollo", "price" : 10, "description" : "Causa Rellena de Pollo" }
{ "_id" : ObjectId("5b144279d7e0df4308b72877"), "productid" : "cep03", "name" : "Curry Extra Picante", "price" : 18, "description" : "curry extra picante" }
{ "_id" : ObjectId("5b144279d7e0df4308b72878"), "productid" : "hch02", "name" : "Honey Chicken", "price" : 12, "description" : "honey chicken" }*/
