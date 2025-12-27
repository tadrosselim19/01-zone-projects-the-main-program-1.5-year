package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,"Hello, world from Go backend!")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the about page.")
    })
	log.Fatal(http.ListenAndServe(":8080",nil))
}