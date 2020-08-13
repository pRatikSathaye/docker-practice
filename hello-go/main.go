package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func start(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"success\"}"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", start).Methods(http.MethodGet)

	fmt.Println("Server started and listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
