package storageUtil

import "go.mongodb.org/mongo-driver/bson/primitive"

type Token struct {
	ID               primitive.ObjectID `bson:"id,omitempty"`
	Access           string             `bson:"access,omitempty"`
	OrganizationCode string             `bson:"organization_code,omitempty"`
}
