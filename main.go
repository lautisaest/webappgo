package main

import (
	"log"
	"net/http"
)

const PortNumer = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	log.Printf("Listening at port %s", PortNumer)
	_ = http.ListenAndServe(PortNumer, nil)

}
