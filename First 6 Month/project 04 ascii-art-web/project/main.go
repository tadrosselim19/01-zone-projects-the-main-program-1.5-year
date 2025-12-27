package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type DataCollect struct {
	Text    string
	Banner  string
	Color   string
	Result  template.HTML
	Massage string
}

var tmpl = template.Must(template.ParseFiles("templates/layout.html"))
var tmpl2 = template.Must(template.ParseFiles("templates/failure.html"))
var text string
var banner string
var color string
var substring string = ""


func main() {

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	// Home page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	// Submit
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {

		text = r.FormValue("text")
		banner = r.FormValue("banner")
		color = r.FormValue("color")
		substring = r.FormValue("substring")

		if text == "" {
			data := DataCollect{
				Massage: "Error YOU must type somthing",
			}
			tmpl2.Execute(w, data)
			return
		}

		// Generate ASCII

		final_result := ""

		if substring != "" {
			array_of_text := split_text_before_coloring(text, substring)
			for _, part := range array_of_text {
				ascii_block := GenerateASCII(part, "banners/"+banner+".txt")
				if part == substring {
					// Wrap only the substring's ASCII block in span with color
					final_result += `<span style="color:` + color + `;">` + ascii_block + `</span>`
				} else {
					final_result += ascii_block
				}
			}
		} else {
			ascii_block := GenerateASCII(text, "banners/"+banner+".txt")
			final_result += `<span style="color:` + color + `;">` + ascii_block + `</span>`
		}

		

		// Mark as safe HTML
		data := DataCollect{
			Text:   text,
			Banner: banner,
			Color:  color,
			Result: template.HTML(final_result),
		}

		tmpl.Execute(w, data)
	})


	// handle downlod of file
	http.HandleFunc("/downlood_as_txt",func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}
		
		if text == "" {
			http.Error(w, "NO Content Provided", http.StatusBadRequest)
			return
		}

		file_result := GenerateASCII(text,"banners/"+ banner +".txt")

		file_result_byte := []byte(file_result)
		w.Header().Set("Content-Type","text/plain")
		w.Header().Set("Content-Length",strconv.Itoa(len(file_result_byte)))
		w.Header().Set("Content-Disposition","attachment; filename=\"ASSCI.txt\"")
		w.Write(file_result_byte)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
