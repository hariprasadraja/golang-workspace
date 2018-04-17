package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	EmailID     string        `json:"emailID" bson:"emailID"`
	DateCreated time.Time     `json:"dateCreated" bson:"dateCreated"`
	LastUpdated time.Time     `json:"lastUpdated" bson:"lastUpdated"`
}
