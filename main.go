package main

import (
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", controller.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", controller.RegisterHandler).Methods("POST")
	r.HandleFunc("/profile", controller.RegisterHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
