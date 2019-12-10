package main

import (
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome to the snippetbox"))
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)

	


}