package main

import (
	"InventoryManagement/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Inventory Management System")
	router := routes.Router()
	log.Fatal(http.ListenAndServe(":8080", router))

}
