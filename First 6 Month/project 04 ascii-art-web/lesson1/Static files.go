// ex3.go
package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("serving ./static at http://localhost:8080/static/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
