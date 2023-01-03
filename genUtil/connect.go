package genUtil

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectToMongo() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(getMongoURI()))
	if err != nil {
		fmt.Println("Error creating client")
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	
	err = client.Connect(ctx)
	//err alrexists, just reassigning
	if err != nil {
		fmt.Println("Error connecting to client")
	}

	fmt.Println("Connected to MDB db")
	return client
}



