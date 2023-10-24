package main

import (
	"fmt"
	"helloworld/internal/routes"

	"net/http"
)

func main() {
	router := routes.NewRouter()
	fmt.Printf("This output")

	fmt.Print("--------------\n")

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
