package storageUtil

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Token struct {
	ID               primitive.ObjectID `bson:"id,omitempty"`
	Access           string             `bson:"access,omitempty"`
	OrganizationCode string             `bson:"organization_code,omitempty"`
	Expiry           time.Time          `bson:"expiry,omitempty"`
}
