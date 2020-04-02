package main

import (
	"fmt"
	"log"
	"net/http"
)

func logRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Printf("%s request %s", req.RemoteAddr, req.URL)
		h.ServeHTTP(res, req)
	})
}

func main() {
	h := http.NewServeMux()

	h.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Home handler")
	})

	h.HandleFunc("/about", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "About handler")
	})

	port := ":5000"

	// logger midleware
	h1 := logRequest(h)

	log.Println("Server running on port", port)

	err := http.ListenAndServe(port, h1)
	if err != nil {
		log.Fatal(err)
	}
}
