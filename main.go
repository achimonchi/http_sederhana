package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server running at port :3000")
	// h := myHandler{}

	// menggunakan handlerFunc dari golang
	h := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello Golang")
		fmt.Println(res)
	})
	err := http.ListenAndServe(":3000", h)
	if err != nil {
		log.Fatal(err)
	}
}

// type myHandler struct {
// }

// // memasukkan myHandler kedalam fungsi ServeHTTP punya golang
// func (m myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(res, "Hello Golang")
// }
