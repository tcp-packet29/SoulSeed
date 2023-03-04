package storageUtil

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
) //represent bson in terms of mongo processing and database handling and parsing, better fop rhandling and manipulating and uitlizing mongo db

type User struct {
	Id                primitive.ObjectID `bson:"id,omitempty"` //bson is a package that allows us to use bson in go, omitempty means that if the field is empty, it will not be included in the bson
	Username          string             `bson:"username,omitempty"`
	Email             string             `bson:"email,omitempty"`
	Password          string             `bson:"password,omitempty"`
	Items             []string           `bson:"items,omitempty"`
	Zipcode           string             `bson:"zipcode,omitempty"`
	OrganizationOwned []string           `bson:"organization_owned,omitempty"`
}

func (usr *User) hashPassword(pass string) error {
	bytearr, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	usr.Password = string(bytearr)
	return err
}

func (usr *User) checkPass(pass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(pass))
	return err
}

//might not use; will implement hashing nd cuberseurity later
