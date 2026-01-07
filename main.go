package main

import (
	"fmt"
	"net/http"
	handler "weather-api/handlers"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.GetWelcome)
	fmt.Println("server run on http://localhost:8383")

	http.ListenAndServe(":8383", mux)
	//server run on http://localhost:8383

}
