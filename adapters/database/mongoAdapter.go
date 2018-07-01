package database

import (
	//"fmt"
	mgo "gopkg.in/mgo.v2"
	"log"
)

var (
	Dao *mgo.Database
)

type MongoDbAdapter struct {
	URL      string
	Database string
}

func (m MongoDbAdapter) Connect() {
	session, err := mgo.Dial(m.URL)
	if err != nil {
		log.Fatal(err)
	}
	Dao = session.DB(m.Database)
}
