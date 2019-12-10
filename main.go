package main

import (
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome to the snippetbox"))
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)

	log.Println("Starting server on : 4000")

	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)

}