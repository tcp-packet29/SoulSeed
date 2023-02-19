package genUtil

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func GetMailgunData() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading environmen file")
	}
	return os.Getenv("MG_APIKEY")

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
