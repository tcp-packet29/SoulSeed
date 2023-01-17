package storageUtil

import "go.mongodb.org/mongo-driver/bson/primitive" //represent bson in terms of mongo processing and database handling and parsing, better fop rhandling and manipulating and uitlizing mongo db

type Organization struct {
	Id 	 primitive.ObjectID `bson:"id,omitempty"` //bson is a package that allows us to use bson in go, omitempty means that if the field is empty, it will not be included in the bson
	Name string `bson:"name,omitempty"`
	Zipcode string `bson:"zipcode,omitempty"`
	Owner User `bson:"owner,omitempty"`
	Owner_ID string `bson:"owner_id"` //conver to tobjectid from hex or string value
	Items []Item `bson:"items,omitempty"`
	Users []User `bson:"users,omitempty"`
	Users_ID []string `bson:"users_id"` //needed to get users from db and ids and convert to obectid form hex or string value can jts reference by id amd upon processing through srever and request handlin will add in user to omitempty fiel
}