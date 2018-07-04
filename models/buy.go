package models

import (
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

type Buy struct {
	ID     bson.ObjectId `bson:"_id" json: "_id"`
	Userid string        `bson:"name" json: "name"`
	Dishes []string      `bson:"dishes" json:"dishes"`
	DoneAt time.Time     `bson:"doneat" json:"doneat"`
	Total  float64       `bson:"total" json:"total"`
}

func (b Buy) String() string {
	out := ""
	out += "UserId: " + b.Userid
	out += "\nDishes:\n"
	for _, v := range b.Dishes {
		out += "\t" + v + "\n"
	}
	out += "DoneAt: " + b.DoneAt.String()
	out += "\nTotal " + strconv.FormatFloat(b.Total, 'f', 4, 64)
	return out
}
