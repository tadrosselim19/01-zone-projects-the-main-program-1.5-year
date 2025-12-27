package main

import (
	"fmt"
	"log"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, "Write a message using POST")
		return
	}

	if r.Method == http.MethodPost {
		fmt.Fprint(w, "Message received!")
		return
	}

	// If any other method is used
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", loginHandler)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
