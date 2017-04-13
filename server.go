package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello")
	router := mux.NewRouter()

	router.HandleFunc("/{image}", HomeHandler)
	http.Handle("/", router)
	http.ListenAndServe(":18080", nil)

}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {

	http.ServeFile(writer, request, "/home/mao/documentations")
	return
}
