package models

import (
	"gopkg.in/mgo.v2/bson"
)

// User struct
type User struct {
	Model

	ID          bson.ObjectId 		`json:"id" bson:"_id"`
	First		string  			`json:"first" bson:"first"`
	Last		string				`json:"last" bson:"last"`
	Email		string				`json:"email" bson:"email"`
}

// GetUsers func - return all users
func GetUsers() ([]User, error) {
	mongo := Session.Copy()
	defer mongo.Close()
	users := []User{}
	err := mongo.DB("yourdb").C("users").Find(nil).All(&users)

	if err != nil {
		return users, err
	}

	return users, nil
}

// GetUserByEmail func
func GetUserByEmail (email string) (User, error){
		
	mongo := Session.Copy()
	defer mongo.Close()
	user := User{}
	err := mongo.DB("yourdb").C("users").Find(bson.M{"email": email}).One(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}

// CreateUser func
func CreateUser (email string, first string, last string) (User, error){
	mongo := Session.Copy()
	defer mongo.Close()

	c := mongo.DB("darwin").C("users")

	user := User{}

	// insert new user record
	id := bson.NewObjectId()
	err := c.Insert(&User{ID: id, First: first, Last: last})

	if err != nil {
		return user, err
	}

	// fetch new record
	err = c.Find(bson.M{"_id": id}).One(&user)

	if err != nil {
		return user, err
	}

	return user, nil

}
