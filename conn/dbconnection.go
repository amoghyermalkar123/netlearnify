package conn

import (
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"os"
)

var db *mgo.Database

func init() {
	host := "localhost"
	dbName := "netldb"

	session, err := mgo.Dial(host)
	if err != nil {
		fmt.Println("session err:", err)
		os.Exit(2)
	}
	db = session.DB(dbName)
}

// GetMongoDB function to return DB connection
func GetMongoDB() *mgo.Database {
	return db
}
