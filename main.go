package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	h := myHandler{}
	err := http.ListenAndServe(":3000", h)
	log.Fatal(err)
}

type myHandler struct {
}

func (m myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello Golang")
}
