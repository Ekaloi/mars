package main

import (
	"fmt"
	"log"

	"github.com/ekaloi/mars/api"
	"github.com/joho/godotenv"
)


func main(){
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	apiClient := api.NewClient()

	data, err := apiClient.GetWeather()
	if err != nil {
		fmt.Printf("Error \n")
	}
	fmt.Printf("%v", data)
}