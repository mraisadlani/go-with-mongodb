package common

import (
	"os"

	"gopkg.in/mgo.v2")


func Connection() (*mgo.Database, error) {
	Host := os.Getenv("MONGO_HOST")
	Name := os.Getenv("MONGO_NAME")

	session, err := mgo.Dial(Host)

	if err != nil {
		panic(err)
	}

	db := session.DB(Name)

	return db, nil
}