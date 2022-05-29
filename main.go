package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, world!")
		if err != nil {
			fmt.Println(err)

		}
		log.Println("Number of ytes written:", n)
	})

	_ = http.ListenAndServe(":8080", nil)

}
