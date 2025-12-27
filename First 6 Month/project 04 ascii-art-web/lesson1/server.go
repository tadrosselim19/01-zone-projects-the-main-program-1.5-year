package main

import (
	"fmt"
	"net/http"
)


func hallohandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"hello, seerver")
}


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",hallohandler)
	err := http.ListenAndServe(":8080",mux)
	if err != nil{
		panic(err)
	}
}