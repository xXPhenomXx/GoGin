package models

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"database/sql"

	"GoGin/pkg/setting"
	"time"
)

var db *sql.DB
var Session *mgo.Session

type Model struct {
	CreatedOn  int `json:"createdOn" bson:"createdOn"`
	ModifiedOn int `json:"modified_on" bson:"modifiedOn"`
	DeletedOn  int `json:"deleted_on" bson:"deletedOn"`
}

func Setup() {
	var err error

	log.Println("Initializing MongoDB connection...")

    defer func() {
        if r := recover(); r != nil {
            log.Fatalf("Mongo panic detected: %v", err)
            var ok bool
            err, ok := r.(error)
            if !ok {
                fmt.Printf("pkg:  %v,  error: %s", r, err)
            }
        }
    }()

	Session,err = mgo.DialWithTimeout(setting.MongoSetting.Host, 60 * time.Second)
	Session.SetSocketTimeout(60 * time.Second)
    if err!=nil{
		log.Println("MongoDB Init Error: %v", err)
        // panic(err)
    }
    log.Println("Successfully connected to MongoDB")
    Session.SetMode(mgo.Monotonic, true)
}

func CloseDB() {
	defer db.Close()
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
