package api

import (
	"autocredit/cmd/api/internal/transport"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var server = transport.Server{}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func Run() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	//seed.Load(server.DB)

	server.Run(":" + os.Getenv("APP_PORT"))
}
