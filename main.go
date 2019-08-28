package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"./src/controller"
)

func main() {
	fmt.Println("Starting the application...")
	r := mux.NewRouter()
	r.HandleFunc("/register", controller.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", controller.LoginHandler).Methods("POST")
	r.HandleFunc("/profile", controller.ProfileHandler).Methods("GET")
	r.HandleFunc("/health", controller.HealthCheck).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Application live and listening at port 8080")
}
