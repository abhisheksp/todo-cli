package crud

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"os"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	list := params["list"]
	host, user, password, database, ssl := os.Getenv("PHOST"), os.Getenv("PUSER"), os.Getenv("PPASSWORD"), os.Getenv("PDB"), os.Getenv("PSSL")
	dataSource := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", host, user, password, database, ssl)
	db, err := sql.Open("postgres", dataSource)

	if err != nil {
		log.Fatalf("Error: The data source arguments are not valid: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error: Could not establish a connection with the database: %v", err)
	}

	queryStmt, err := db.Prepare("SELECT item FROM todo WHERE list = $1")
	var items []string
	rows, err := queryStmt.Query(list)
	defer rows.Close()

	for rows.Next() {
		var item string
		if err := rows.Scan(&item); err != nil {
			log.Fatal(err)
		}
		items = append(items, item)
	}

	fmt.Fprintf(w, "%s\n", list)
	fmt.Fprintln(w, "-------------------------------------")
	fmt.Fprintf(w, "%v", items)
}
