package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"helloworld/internal/routes"
	"log"
	"net/http"
)

type MyData struct {
	ID        int
	Migration string
	Batch     int
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/trials-database")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, migration, batch FROM migrations")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var data []MyData

	for rows.Next() {
		var d MyData
		err := rows.Scan(&d.ID, &d.Migration, &d.Batch)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, d)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

	fmt.Printf("This output\n")

	fmt.Print("--------------\n")

	newServer(data)

}

func newServer(data []MyData) {

	router := routes.NewRouter(data)

	port := 8080
	address := fmt.Sprintf(":%d", port)
	fmt.Printf("Server is listening on http://localhost%s\n", address)
	err := http.ListenAndServe(
		address,
		router,
	)
	if err != nil {
		panic(err)
	}
}
