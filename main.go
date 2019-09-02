package main

import (
	"fmt"

	"./src/controller"
)

func main() {
	fmt.Println("Starting the application...")

	db := db.ConnectToDatabase

	c := controller.StartController

	r := router.Router(c)

	r.StartWithCORS(r)

	fmt.Println("Application live and listening at port 8080")
}
