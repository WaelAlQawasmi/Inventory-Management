package routes

import (
	"InventoryManagement/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/items", controllers.CreateItem).Methods("POST")
	router.HandleFunc("/items", controllers.GetItem).Methods("GET")
	return router
}
