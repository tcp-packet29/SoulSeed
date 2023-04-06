package dbUtil

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

// single instance of mongo client so we dont createa new client t o access the server from every time
var client *mongo.Client = ConnectToMongo()

func GetClient() *mongo.Client {
	return client
}

func GetCollection(collectionName string) *mongo.Collection {
	return client.Database("fruitApp").Collection(collectionName)
}
func GetGridfs() *mongo.Database {
	return client.Database("gridfs")
}

func GetGridfsBucket() *gridfs.Bucket {
	buclet, err := gridfs.NewBucket(GetGridfs())
	//idempotence
	if err != nil {
		panic(err)
	}
	return buclet
}
