package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id           bson.ObjectId `json:"id" bson:"_id"`
	FirstName    string        `json:"name" bson:"name"`
	LastName     string        `json:"lastname" bson:"lastname"`
	UserName     string        `json:"username" bson:"username"`
	OrgName      string        `json:"orgname" bson:"orgname"`
	Address1     string        `json:"address1" bson:"address1"`
	Address2     string        `json:"address2" bson:"address2"`
	City         string        `json:"city" bson:"city"`
	State        string        `json:"State" bson:"State"`
	ZipCode      string        `json:"zip" bson:"zip"`
	Country      string        `json:"country" bson:"country"`
	Gender       string        `json:"gender" bson:"gender"`
	Account      string        `json:"account" bson:"account"`
	DOB          time.Time     `json:"dob" bson:"dob"`
	Email        string        `json:"email" bson:"email"`
	Password     string        `json:"password" bson:"password"`
	Active       bool          `json:"active" bson:"active"`
	LastAccessed time.Time     `json:"lastaccessed" bson:"lastaccessed"`
	CreatedDate  time.Time     `json:"createddate" bson:"createddate"`
}
