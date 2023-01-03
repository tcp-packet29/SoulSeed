package dbUtil

import (
	"main/genUtil"
	"go.mongodb.org/mongo-driver/mongo"
	
)


//single instance of mongo client so we dont createa new client t o access the server from every time
var client *mongo.Client = genUtil.ConnectToMongo()

func GetClient() *mongo.Client {
	return client
}

func GetCollection(collectionName string) *mongo.Collection {
	return client.Database("Cluster0").Collection(collectionName)
}

