package storageUtil

import "go.mongodb.org/mongo-driver/bson/primitive" //bson for mongodb manipuklationa nd naviagtion as awelkla s procerssing an difnindgrepresent bson in terms of mongo processing and database handling and parsing, better fop rhandling and manipulating and uitlizing mongo db

type Trade struct {
	Id          primitive.ObjectID `bson:"id,omitempty"`
	Maker       User               `bson:"maker,omitempty"`
	MakerID     string             `bson:"makerid"`
	Name        string             `bson:"name"`
	Items       []string           `bson:"items"` //cooiuild make an ite3m struc tso that they use the same id and can chacek if have nase don id but im usign stribngs for now
	Description string             `bson:"desc"`
	Open        bool               `bson:"open,omitempty"`
	OrgId       string             `bson:"orgid"`
	Latlong     []string           `bson:"latlong"`
	Offers      []Offer            `bson:"offers"`
	Providing   bool               `bson:"providing"`
}

type Model struct {
	OrganizationCode string   `json:"organization_code"`
	Emails           []string `json:"emails"`
}

type Data struct {
	Id  string         `bson:"id"`
	Map map[string]int `bson:"map"`
}

type Offer struct {
	UserId      string `bson:"userid"`
	Usr         User   `bson:"user"`
	Address     string `bson:"address"`
	Description string `bson:"memo"`
	Approved    bool   `bson:"approved"`
}

type Mess struct {
	Message string `json:"message"`
}
