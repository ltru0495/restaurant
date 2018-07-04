package models

import (
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type Dish struct {
	ID          bson.ObjectId `bson:"_id" json: "_id"`
	Dishid      string        `bson: "dishid json:"dishid"`
	Name        string        `bson:"name" json: "name"`
	Description string        `bson:"description" json: "description"`
	Price       float64       `bson:"price" json: "price"`
	Image       string        `bson:"image" json: "image"`
}

func (d Dish) String() string {
	return "\nDishId:" + d.Dishid + "\nName: " + d.Name + "\nDescription: " +
		d.Description + "\nPrice: " + strconv.FormatFloat(d.Price, 'f', 4, 64) + "\nImage:" + d.Image + "\n"
}
