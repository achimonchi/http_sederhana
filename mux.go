package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// membuat server dengan mux
	// saat membuat NewServeMux(), maka sudah langsung mengimplementasi
	// http.Handler
	h := http.NewServeMux()

	// membuat route
	h.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Home handler")
	})
	h.HandleFunc("/about", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "About handler")
	})

	port := ":4000"

	log.Println("Server running on port", port)

	err := http.ListenAndServe(port, h)
	if err != nil {
		log.Fatal(err)
	}
}
