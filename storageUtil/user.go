package storageUtil

import "go.mongodb.org/mongo-driver/bson/primitive" //represent bson in terms of mongo processing and database handling and parsing, better fop rhandling and manipulating and uitlizing mongo db

type User struct {
	Id 	 primitive.ObjectID `bson:"id,omitempty"` //bson is a package that allows us to use bson in go, omitempty means that if the field is empty, it will not be included in the bson
	Username string `bson:"username,omitempty"`
	Password string `bson:"password,omitempty"`
	Items []string `bson:"items,omitempty"`
	Zipcode string `bson:"zipcode,omitempty"`
}