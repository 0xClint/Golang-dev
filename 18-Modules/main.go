package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	greeter()
}

func greeter() {
	fmt.Println("Modules in GOlang")

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	// w - Write, the response that u will be writing
	// r - Readed, the response that u will be reading whoever is sending to you
	w.Write([]byte("<h1>Namaste World!</h1>"))
}
