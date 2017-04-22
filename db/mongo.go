package db

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

func Init(host, dbname string) (*mgo.Database, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		log.Println("can not conect mongodb")
		return nil, err
	}

	database := session.DB(dbname)
	return database, nil
}
