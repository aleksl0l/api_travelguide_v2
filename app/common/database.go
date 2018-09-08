package common

import (
	"api_travelguide_v2/app"
	"fmt"
	"github.com/globalsign/mgo"
)

type Database struct {
	*mgo.Database
}

var DB *mgo.Database

func Init() *mgo.Database {
	session, err := mgo.Dial(app.DB_URI)
	if err != nil {
		fmt.Println("db error: ", err)
	}
	DB = session.DB(app.DB_NAME)
	return DB
}

func GetDb() *mgo.Database {
	return DB
}
