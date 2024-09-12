package main

import "github.com/joho/godotenv"

import (
	"github.com/rohit1kumar/pgo/config"
	"github.com/rohit1kumar/pgo/models"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectToDB()
}

func main() {
	err := config.DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal(err)
	}
}
