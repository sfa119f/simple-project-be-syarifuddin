package main

import (
	"fmt"
	"net/http"

	"simple-project-be/backend/database"
	"simple-project-be/backend/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// init db
	database.InitDB()

	// init router
	router := mux.NewRouter()

	// routes
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	}).Methods(http.MethodGet)
	router.HandleFunc("/products", handlers.GetProducts).Methods(http.MethodGet)

	http.ListenAndServe(":8000", router)
}
