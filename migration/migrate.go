package main

import (
	"github.com/joho/godotenv"
	"github.com/rohit1kumar/pgo/config"
	"github.com/rohit1kumar/pgo/models"
	"log"
)

func init() {
	godotenv.Load()
	config.ConnectToDB()
}

func main() {
	err := config.DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal(err)
	}
}
