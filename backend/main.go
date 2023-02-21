package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// init router
	router := mux.NewRouter()

	// routes
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	}).Methods(http.MethodGet)

	http.ListenAndServe(":8000", router)
}
