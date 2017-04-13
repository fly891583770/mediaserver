package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func YourHandler2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla222!\n"))
}

func main() {
	r := mux.NewRouter()
	r = r.Path("/test").Subrouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/1", YourHandler)
	r.HandleFunc("/2", YourHandler2)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":18080", r))
}
