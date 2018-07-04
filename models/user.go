package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId `bson: "_id,omitempty" json: "id, omitempty"`
	Userid   string        `bson:"userid" json:"userid"`
	Name     string        `bson:"name" json: "name"`
	Surname  string        `bson:"surname" json: "surname"`
	Phone    string        `bson:"phone" json:"phone"`
	Email    string        `bson:"email" json: "email"`
	Username string        `bson:"username" json: "username"`
	Password string        `bson:"password" json: "password"`
	Type     string        `bson:"type" json:"type"`
}

func (u User) String() string {
	return "Name: " + u.Name + " " + u.Surname +
		"\nPhone: " + u.Phone + "\nEmail: " + u.Email +
		"\nUsername: " + u.Username + "UserId: " + u.Userid + "\n"
}
