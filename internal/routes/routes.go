package routes

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", apiDataHandler)

	return mux
}

func apiDataHandler(responseWriter http.ResponseWriter, request *http.Request) {
	data := "some api Data"
	_, err := fmt.Fprintln(responseWriter, data)
	if err != nil {
		panic(err)
	}
}
