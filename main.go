package main

import (
	// "fmt"
	// "os"
	// "context"
	"github.com/gin-gonic/gin"
	"main/genUtil"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	genUtil.ConnectToMongo()
	//mnot actually an instance of client we can actually reference, just for tersrting purposes
	router.Run(":8080")
}