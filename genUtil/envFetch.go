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

func GetGoogleData() (string, string) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv("GOOGLE_CLIENT"), os.Getenv("GOOGLE_SECRET")
}

func GetJWTData() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv("JWT_SECRET")
}
