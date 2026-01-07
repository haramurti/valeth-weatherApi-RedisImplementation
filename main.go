package main

import (
	"fmt"
	"net/http"
	handler "weather-api/handlers"

	"github.com/joho/godotenv"
)

func main() {

	mux := http.NewServeMux()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("cannot load env")
	}
	fmt.Println("env loaded..")

	mux.HandleFunc("/", handler.GetWelcome)
	mux.HandleFunc("/api/v1/weather/{city}", handler.GetCityWeather)

	fmt.Println("server run on http://localhost:8383")
	http.ListenAndServe(":8383", mux)
	//server run on http://localhost:8383

}
