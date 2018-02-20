package main

import (
	"log"
	"net/http"
	"os"

	"github.com/abhisheksp/todo-cli/crud"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := mux.NewRouter()
	router.HandleFunc("/todo/{list}/{item}", crud.CreateItem).Methods("POST")
	router.HandleFunc("/todo/{list}", crud.GetItems).Methods("GET")
	log.Println("Running at :8000")
	log.Fatal(http.ListenAndServe(":" + port, router))
}
