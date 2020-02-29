package main

import (
	"controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/", controllers.Index).Methods("GET")
	router.HandleFunc("/upload", controllers.Upload).Methods("POST")

	http.Handle("/", router)
	fmt.Println("working on :8080")
	http.ListenAndServe(":8080", nil)
}
