package main

import (
	"log"
	"net/http"

	"github.com/abhisheksp/todo-cli/crud"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/todo/{list}/{item}", crud.CreateItem).Methods("POST")
	router.HandleFunc("/todo/{list}", crud.GetItems).Methods("GET")
	log.Println("Running at :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
