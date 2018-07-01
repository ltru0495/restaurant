package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Menu struct {
	ID     bson.ObjectId `bson:"_id" json: "_id"`
	Dishes []string      `bson:"dishes" json:"dishes"`
	Date   time.Time     `bson:"date" json:date`
}

func (m Menu) String() string {
	out := ""
	out += "\nDishes:\n"
	for _, v := range m.Dishes {
		out += "\t" + v + "\n"
	}
	out += m.Date.String()
	return out
}
