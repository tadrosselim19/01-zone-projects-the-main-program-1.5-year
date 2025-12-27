package main

import (
	"fmt"
	"log"
	"net/http"
)

func handle_great(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodGet{
		fmt.Fprintf(w,"Method Not Allowed")
		return
	}
	if r.URL.Query().Get("name")== ""{
		fmt.Fprintf(w,"hello, stranger!")
		return
	}
	fmt.Fprintf(w,"hello, %s!",r.URL.Query().Get("name"))
}
func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/great",handle_great)

	log.Fatal(http.ListenAndServe(":8080",mux))
	
}