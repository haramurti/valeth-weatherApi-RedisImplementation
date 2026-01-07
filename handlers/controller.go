package handler

import "net/http"

func GetWelcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to weather api"))
}
