package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type page_data struct {
	Name  string
	Items []string
}
var html1 = template.Must(template.ParseFiles("01 talent projects/frist 6 months/project 04 ascii-art-web/lesson2/templates/hallo.html"))
		// t, err := template.ParseFiles("templates/index.html")
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
		// if err != nil {
		// 	panic(err)
		// }

		// data := map[string]string{
		// 	"name": "tadros",
		// }
		data := page_data{
			Name:  "tadros",
			Items: []string{"Go", "Python", "Rust"},
		}
		html1.Execute(w, data)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		html := template.Must(template.ParseFiles("01 talent projects/frist 6 months/project 04 ascii-art-web/lesson2/templates/login.html"))
		if r.Method == http.MethodPost {
			name := r.FormValue("name")
			data := page_data{
				Name: name,
			}
			html.Execute(w, data)
			return
		}else{
			fmt.Fprintf(w,"method not allowed")
			html1.Execute(w,nil)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
