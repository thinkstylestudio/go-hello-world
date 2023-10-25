package routes

import (
	"fmt"
	"helloworld"
	"net/http"
)

var apiData string

func NewRouter(data []main.MyData) http.Handler {
	apiData = data
	mux := http.NewServeMux()
	mux.HandleFunc("/", apiDataHandler)

	return mux
}

//const apiData = "some api Data"

func apiDataHandler(responseWriter http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintln(responseWriter, apiData)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}
