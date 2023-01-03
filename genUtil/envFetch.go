package genUtil

import (
	"github.com/joho/godotenv"
	"fmt"
	"os"
)

func getMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv("MURI")
}
