package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

// Structure to pass data to template
type PageData struct {
	Text   string
	Banner string
	Result string
	Error  string
}

// ✅ Function to load a banner file
func loadBanner(banner string) (map[rune]string, error) {
	filePath := fmt.Sprintf("banners/%s.txt", banner)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read banner: %v", err)
	}

	lines := strings.Split(string(content), "\n")
	asciiMap := make(map[rune]string)
	r := rune(32) // printable ASCII start

	for i := 0; i < len(lines); i += 9 {
		if r > 126 {
			break
		}
		charLines := strings.Join(lines[i:i+8], "\n")
		asciiMap[r] = charLines
		r++
	}

	return asciiMap, nil
}

// ✅ Function to generate ASCII art
func generateASCII(text, banner string) (string, error) {
	if text == "" {
		return "", fmt.Errorf("empty text")
	}

	bannerMap, err := loadBanner(banner)
	if err != nil {
		return "", err
	}

	lines := strings.Split(text, "\\n")
	var result strings.Builder

	for _, line := range lines {
		for row := 0; row < 8; row++ {
			for _, char := range line {
				ascii, ok := bannerMap[char]
				if !ok {
					return "", fmt.Errorf("character %q not found in banner", char)
				}
				asciiLines := strings.Split(ascii, "\n")
				result.WriteString(asciiLines[row])
			}
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}

// ✅ Handler for GET /
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl := template.Must(template.ParseFiles("index2.html"))
	data := PageData{}
	tmpl.Execute(w, data)
}

// ✅ Handler for POST /ascii-art
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	result, err := generateASCII(text, banner)
	data := PageData{Text: text, Banner: banner}

	if err != nil {
		data.Error = err.Error()
	} else {
		data.Result = result
	}

	tmpl := template.Must(template.ParseFiles("index2.html"))
	tmpl.Execute(w, data)
}

// ✅ Main function
func main() {
	fmt.Println("Starting server on http://localhost:8080")
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
