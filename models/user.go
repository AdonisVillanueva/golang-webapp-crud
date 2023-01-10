package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Gender   string        `json:"gender" bson:"gender"`
	Account  string        `json:"account" bson:"account"`
	DOB      time.Time     `json:"dob" bson:"dob"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"password" bson:"password"`
	Active   bool          `json:"active" bson:"active"`
}
