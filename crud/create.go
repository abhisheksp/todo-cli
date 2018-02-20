package crud

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	list, item := params["list"], params["item"]

	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("Error: The data source arguments are not valid: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error: Could not establish a connection with the database: %v", err)
	}

	createQuery := "INSERT INTO todo(list, item) VALUES($1, $2)"

	_, err = db.Exec(createQuery, list, item)
	if err != nil {
		log.Fatalf("Error: Could not insert item: %v", err)
	}

	fmt.Fprintf(w, "added \"%s\" to the \"%s\" list", item, list)
}
