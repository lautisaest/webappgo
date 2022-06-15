package main

import (
	"log"
	"net/http"

	handlers "github.com/lautisaest/webappgo/pkg/handlers"
)

const PortNumer = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	log.Printf("Listening at port %s", PortNumer)
	_ = http.ListenAndServe(PortNumer, nil)

}
