package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	EmailID     string        `json:"emailID" bson:"emailID"`
	DateCreated time.Time     `json:"dateCreated" bson:"dateCreated"`
	LastUpdated time.Time     `json:"lastUpdated" bson:"lastUpdated"`
}
